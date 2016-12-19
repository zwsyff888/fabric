/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package endorser

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/op/go-logging"
	"golang.org/x/net/context"

	"github.com/hyperledger/fabric/core/chaincode"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	putils "github.com/hyperledger/fabric/protos/utils"
)

var endorserLogger = logging.MustGetLogger("endorser")

// The Jira issue that documents Endorser flow along with its relationship to
// the lifecycle chaincode - https://jira.hyperledger.org/browse/FAB-181

// Endorser provides the Endorser service ProcessProposal
type Endorser struct {
}

// NewEndorserServer creates and returns a new Endorser server instance.
func NewEndorserServer() pb.EndorserServer {
	e := new(Endorser)
	return e
}

//TODO - what would Endorser's ACL be ?
func (*Endorser) checkACL(prop *pb.Proposal) error {
	return nil
}

//TODO - check for escc and vscc
func (*Endorser) checkEsccAndVscc(prop *pb.Proposal) error {
	return nil
}

func (*Endorser) getTxSimulator(ledgername string) (ledger.TxSimulator, error) {
	lgr := peer.GetLedger(ledgername)
	if lgr == nil {
		return nil, fmt.Errorf("chain does not exist(%s)", ledgername)
	}
	return lgr.NewTxSimulator()
}

//deploy the chaincode after call to the system chaincode is successful
func (e *Endorser) deploy(ctxt context.Context, cccid *chaincode.CCContext, cds *pb.ChaincodeDeploymentSpec) error {
	chaincodeSupport := chaincode.GetChain()

	_, err := chaincodeSupport.Deploy(ctxt, cccid, cds)
	if err != nil {
		return fmt.Errorf("Failed to deploy chaincode spec(%s)", err)
	}

	//launch and wait for ready
	_, _, err = chaincodeSupport.Launch(ctxt, cccid, cds)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	//stop now that we are done
	chaincodeSupport.Stop(ctxt, cccid, cds)

	return nil
}

//call specified chaincode (system or user)
func (e *Endorser) callChaincode(ctxt context.Context, chainID string, txid string, prop *pb.Proposal, cis *pb.ChaincodeInvocationSpec, cid *pb.ChaincodeID, txsim ledger.TxSimulator) ([]byte, *pb.ChaincodeEvent, error) {
	var err error
	var b []byte
	var ccevent *pb.ChaincodeEvent

	if txsim != nil {
		ctxt = context.WithValue(ctxt, chaincode.TXSimulatorKey, txsim)
	}

	//is this a system chaincode
	syscc := chaincode.IsSysCC(cid.Name)

	cccid := chaincode.NewCCContext(chainID, cid.Name, "", txid, syscc, prop)

	b, ccevent, err = chaincode.ExecuteChaincode(ctxt, cccid, cis.ChaincodeSpec.CtorMsg.Args)

	if err != nil {
		return nil, nil, err
	}

	//----- BEGIN -  SECTION THAT MAY NEED TO BE DONE IN LCCC ------
	//if this a call to deploy a chaincode, We need a mechanism
	//to pass TxSimulator into LCCC. Till that is worked out this
	//special code does the actual deploy, upgrade here so as to collect
	//all state under one TxSimulator
	//
	//NOTE that if there's an error all simulation, including the chaincode
	//table changes in lccc will be thrown away
	if cid.Name == "lccc" && len(cis.ChaincodeSpec.CtorMsg.Args) == 3 && (string(cis.ChaincodeSpec.CtorMsg.Args[0]) == "deploy" || string(cis.ChaincodeSpec.CtorMsg.Args[0]) == "upgrade") {
		var cds *pb.ChaincodeDeploymentSpec
		cds, err = putils.GetChaincodeDeploymentSpec(cis.ChaincodeSpec.CtorMsg.Args[2])
		if err != nil {
			return nil, nil, err
		}

		//this should not be a system chaincode
		if chaincode.IsSysCC(cds.ChaincodeSpec.ChaincodeID.Name) {
			return nil, nil, fmt.Errorf("attempting to deploy a system chaincode %s/%s", cds.ChaincodeSpec.ChaincodeID.Name, chainID)
		}

		cccid = chaincode.NewCCContext(chainID, cds.ChaincodeSpec.ChaincodeID.Name, "", txid, false, prop)

		err = e.deploy(ctxt, cccid, cds)
		if err != nil {
			return nil, nil, err
		}
	}
	//----- END -------

	return b, ccevent, err
}

//simulate the proposal by calling the chaincode
func (e *Endorser) simulateProposal(ctx context.Context, chainID string, txid string, prop *pb.Proposal, cid *pb.ChaincodeID, txsim ledger.TxSimulator) ([]byte, []byte, *pb.ChaincodeEvent, error) {
	//we do expect the payload to be a ChaincodeInvocationSpec
	//if we are supporting other payloads in future, this be glaringly point
	//as something that should change
	cis, err := putils.GetChaincodeInvocationSpec(prop)
	if err != nil {
		return nil, nil, nil, err
	}
	//---1. check ACL
	if err = e.checkACL(prop); err != nil {
		return nil, nil, nil, err
	}

	//---2. check ESCC and VSCC for the chaincode
	if err = e.checkEsccAndVscc(prop); err != nil {
		return nil, nil, nil, err
	}

	//---3. execute the proposal and get simulation results
	var simResult []byte
	var resp []byte
	var ccevent *pb.ChaincodeEvent
	resp, ccevent, err = e.callChaincode(ctx, chainID, txid, prop, cis, cid, txsim)
	if err != nil {
		return nil, nil, nil, err
	}

	if txsim != nil {
		if simResult, err = txsim.GetTxSimulationResults(); err != nil {
			return nil, nil, nil, err
		}
	}

	return resp, simResult, ccevent, nil
}

func (e *Endorser) getCDSFromLCCC(ctx context.Context, chainID string, txid string, prop *pb.Proposal, chaincodeID string, txsim ledger.TxSimulator) ([]byte, error) {
	ctxt := ctx
	if txsim != nil {
		ctxt = context.WithValue(ctx, chaincode.TXSimulatorKey, txsim)
	}

	return chaincode.GetCDSFromLCCC(ctxt, txid, prop, chainID, chaincodeID)
}

//endorse the proposal by calling the ESCC
func (e *Endorser) endorseProposal(ctx context.Context, chainID string, txid string, proposal *pb.Proposal, simRes []byte, event *pb.ChaincodeEvent, visibility []byte, ccid *pb.ChaincodeID, txsim ledger.TxSimulator) (*pb.ProposalResponse, error) {
	endorserLogger.Infof("endorseProposal starts for chainID %s, ccid %s", chainID, ccid)

	// 1) extract the chaincodeDeploymentSpec for the chaincode we are invoking; we need it to get the escc
	var escc string
	if !chaincode.IsSysCC(ccid.Name) {
		depPayload, err := e.getCDSFromLCCC(ctx, chainID, txid, proposal, ccid.Name, txsim)
		if err != nil {
			return nil, fmt.Errorf("failed to obtain cds for %s - %s", ccid, err)
		}

		_, err = putils.GetChaincodeDeploymentSpec(depPayload)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal cds for %s - %s", ccid, err)
		}

		// FIXME: pick the right escc from cds - currently cds doesn't have this info
		escc = "escc"
	} else {
		// FIXME: getCDSFromLCCC seems to fail for lccc - not sure this is expected?
		escc = "escc"
	}

	endorserLogger.Infof("endorseProposal info: escc for cid %s is %s", ccid, escc)

	// marshalling event bytes
	var err error
	var eventBytes []byte
	if event != nil {
		eventBytes, err = putils.GetBytesChaincodeEvent(event)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal event bytes - %s", err)
		}
	}

	// 3) call the ESCC we've identified
	// arguments:
	// args[0] - function name (not used now)
	// args[1] - serialized Header object
	// args[2] - serialized ChaincodeProposalPayload object
	// args[3] - binary blob of simulation results
	// args[4] - serialized events
	// args[5] - payloadVisibility
	args := [][]byte{[]byte(""), proposal.Header, proposal.Payload, simRes, eventBytes, visibility}
	ecccis := &pb.ChaincodeInvocationSpec{ChaincodeSpec: &pb.ChaincodeSpec{Type: pb.ChaincodeSpec_GOLANG, ChaincodeID: &pb.ChaincodeID{Name: escc}, CtorMsg: &pb.ChaincodeInput{Args: args}}}
	prBytes, _, err := e.callChaincode(ctx, chainID, txid, proposal, ecccis, &pb.ChaincodeID{Name: escc}, txsim)
	if err != nil {
		return nil, err
	}

	// Note that we do not extract any simulation results from
	// the call to ESCC. This is intentional becuse ESCC is meant
	// to endorse (i.e. sign) the simulation results of a chaincode,
	// but it can't obviously sign its own. Furthermore, ESCC runs
	// on private input (its own signing key) and so if it were to
	// produce simulationr results, they are likely to be different
	// from other ESCCs, which would stand in the way of the
	// endorsement process.

	//3 -- respond
	pResp, err := putils.GetProposalResponse(prBytes)
	if err != nil {
		return nil, err
	}

	return pResp, nil
}

// ProcessProposal process the Proposal
func (e *Endorser) ProcessProposal(ctx context.Context, signedProp *pb.SignedProposal) (*pb.ProposalResponse, error) {
	// at first, we check whether the message is valid
	prop, _, hdrExt, err := peer.ValidateProposalMessage(signedProp)
	if err != nil {
		return &pb.ProposalResponse{Response: &pb.Response{Status: 500, Message: err.Error()}}, err
	}

	hdr, err := putils.GetHeader(prop.Header)
	if err != nil {
		return &pb.ProposalResponse{Response: &pb.Response{Status: 500, Message: err.Error()}}, err
	}

	chainID := hdr.ChainHeader.ChainID

	//chainless MSPs have "" chain name
	ischainless := chaincode.IsChainlessSysCC(hdrExt.ChaincodeID.Name)

	//chainID should be empty for chainless SysCC (such as CSCC for Join proposal) and for
	//nothing else
	if chainID == "" && !ischainless {
		err = fmt.Errorf("chainID not provided for chaincode %s", hdrExt.ChaincodeID.Name)
		return &pb.ProposalResponse{Response: &pb.Response{Status: 500, Message: err.Error()}}, err
	} else if chainID != "" && ischainless {
		err = fmt.Errorf("chainID %s provided for a chainless syscc", hdrExt.ChaincodeID.Name)
		return &pb.ProposalResponse{Response: &pb.Response{Status: 500, Message: err.Error()}}, err
	}

	//TODO check for uniqueness of prop.TxID with ledger

	txid := hdr.ChainHeader.TxID
	if txid == "" {
		err = fmt.Errorf("Invalid txID")
		return &pb.ProposalResponse{Response: &pb.Response{Status: 500, Message: err.Error()}}, err
	}

	// obtaining once the tx simulator for this proposal. This will be nil
	// for chainless proposals
	var txsim ledger.TxSimulator
	if chainID != "" {
		if txsim, err = e.getTxSimulator(chainID); err != nil {
			return &pb.ProposalResponse{Response: &pb.Response{Status: 500, Message: err.Error()}}, err
		}
		defer txsim.Done()
	}
	//this could be a request to a chainless SysCC

	// TODO: if the proposal has an extension, it will be of type ChaincodeAction;
	//       if it's present it means that no simulation is to be performed because
	//       we're trying to emulate a submitting peer. On the other hand, we need
	//       to validate the supplied action before endorsing it

	//1 -- simulate
	//TODO what do we do with response ? We need it for Invoke responses for sure
	//Which field in PayloadResponse will carry return value ?
	result, simulationResult, ccevent, err := e.simulateProposal(ctx, chainID, txid, prop, hdrExt.ChaincodeID, txsim)
	if err != nil {
		return &pb.ProposalResponse{Response: &pb.Response{Status: 500, Message: err.Error()}}, err
	}

	//2 -- endorse and get a marshalled ProposalResponse message
	var pResp *pb.ProposalResponse

	//TODO till we implement global ESCC, CSCC for system chaincodes
	//chainless proposals (such as CSCC) don't have to be endorsed
	if ischainless {
		pResp = &pb.ProposalResponse{Response: &pb.Response{}}
	} else {
		pResp, err = e.endorseProposal(ctx, chainID, txid, prop, simulationResult, ccevent, hdrExt.PayloadVisibility, hdrExt.ChaincodeID, txsim)
		if err != nil {
			return &pb.ProposalResponse{Response: &pb.Response{Status: 500, Message: err.Error()}}, err
		}
	}

	//TODO what do we do with response ? We need it for Invoke responses for sure
	// Set the proposal response payload - it
	// contains the "return value" from the
	// chaincode invocation
	pResp.Response.Payload = result

	return pResp, nil
}

// Only exposed for testing purposes - commit the tx simulation so that
// a deploy transaction is persisted and that chaincode can be invoked.
// This makes the endorser test self-sufficient
func (e *Endorser) commitTxSimulation(proposal *pb.Proposal, chainID string, signer msp.SigningIdentity, pResp *pb.ProposalResponse) error {
	tx, err := putils.CreateSignedTx(proposal, signer, pResp)
	if err != nil {
		return err
	}

	lgr := peer.GetLedger(chainID)
	if lgr == nil {
		return fmt.Errorf("failure while looking up the ledger")
	}

	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return err
	}
	block := common.NewBlock(1, []byte{})
	block.Data.Data = [][]byte{txBytes}
	block.Header.DataHash = block.Data.Hash()
	if err = lgr.Commit(block); err != nil {
		return err
	}

	return nil
}
