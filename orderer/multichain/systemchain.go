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

package multichain

import (
	"fmt"

	"github.com/hyperledger/fabric/common/configtx"
	configvaluesapi "github.com/hyperledger/fabric/common/configvalues"
	configtxorderer "github.com/hyperledger/fabric/common/configvalues/channel/orderer"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/orderer/common/filter"
	cb "github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
	"github.com/hyperledger/fabric/protos/utils"

	"github.com/golang/protobuf/proto"
)

// Define some internal interfaces for easier mocking
type chainCreator interface {
	newChain(configTx *cb.Envelope)
}

type limitedSupport interface {
	PolicyManager() policies.Manager
	SharedConfig() configvaluesapi.Orderer
}

type systemChainCommitter struct {
	filter   *systemChainFilter
	configTx *cb.Envelope
}

func (scc *systemChainCommitter) Isolated() bool {
	return true
}

func (scc *systemChainCommitter) Commit() {
	scc.filter.cc.newChain(scc.configTx)
}

type systemChainFilter struct {
	cc      chainCreator
	support limitedSupport
}

func newSystemChainFilter(ls limitedSupport, cc chainCreator) filter.Rule {
	return &systemChainFilter{
		support: ls,
		cc:      cc,
	}
}

func (scf *systemChainFilter) Apply(env *cb.Envelope) (filter.Action, filter.Committer) {
	msgData := &cb.Payload{}

	err := proto.Unmarshal(env.Payload, msgData)
	if err != nil {
		return filter.Forward, nil
	}

	if msgData.Header == nil /* || msgData.Header.ChannelHeader == nil */ {
		return filter.Forward, nil
	}

	chdr, err := utils.UnmarshalChannelHeader(msgData.Header.ChannelHeader)
	if err != nil {
		return filter.Forward, nil
	}

	if chdr.Type != int32(cb.HeaderType_ORDERER_TRANSACTION) {
		return filter.Forward, nil
	}

	configTx := &cb.Envelope{}
	err = proto.Unmarshal(msgData.Data, configTx)
	if err != nil {
		return filter.Reject, nil
	}

	err = scf.authorizeAndInspect(configTx)
	if err != nil {
		logger.Debugf("Rejecting channel creation because: %s", err)
		return filter.Reject, nil
	}

	return filter.Accept, &systemChainCommitter{
		filter:   scf,
		configTx: configTx,
	}
}

func (scf *systemChainFilter) authorize(configEnvelope *cb.ConfigEnvelope) error {
	if configEnvelope.LastUpdate == nil {
		return fmt.Errorf("Must include a config update")
	}

	configEnvEnvPayload, err := utils.UnmarshalPayload(configEnvelope.LastUpdate.Payload)
	if err != nil {
		return fmt.Errorf("Failing to validate chain creation because of payload unmarshaling error: %s", err)
	}

	configUpdateEnv, err := configtx.UnmarshalConfigUpdateEnvelope(configEnvEnvPayload.Data)
	if err != nil {
		return fmt.Errorf("Failing to validate chain creation because of config update envelope unmarshaling error: %s", err)
	}

	config, err := configtx.UnmarshalConfigUpdate(configUpdateEnv.ConfigUpdate)
	if err != nil {
		return fmt.Errorf("Failing to validate chain creation because of unmarshaling error: %s", err)
	}

	if config.WriteSet == nil {
		return fmt.Errorf("Failing to validate channel creation because WriteSet is nil")
	}

	ordererGroup, ok := config.WriteSet.Groups[configtxorderer.GroupKey]
	if !ok {
		return fmt.Errorf("Rejecting channel creation because it is missing orderer group")
	}

	creationConfigItem, ok := ordererGroup.Values[configtx.CreationPolicyKey]
	if !ok {
		return fmt.Errorf("Failing to validate chain creation because no creation policy included")
	}

	creationPolicy := &ab.CreationPolicy{}
	err = proto.Unmarshal(creationConfigItem.Value, creationPolicy)
	if err != nil {
		return fmt.Errorf("Failing to validate chain creation because first config item could not unmarshal to a CreationPolicy: %s", err)
	}

	ok = false
	for _, chainCreatorPolicy := range scf.support.SharedConfig().ChainCreationPolicyNames() {
		if chainCreatorPolicy == creationPolicy.Policy {
			ok = true
			break
		}
	}

	if !ok {
		return fmt.Errorf("Failed to validate chain creation because chain creation policy (%s) is not authorized for chain creation", creationPolicy.Policy)
	}

	policy, ok := scf.support.PolicyManager().GetPolicy(creationPolicy.Policy)
	if !ok {
		return fmt.Errorf("Failed to get policy for chain creation despite it being listed as an authorized policy")
	}

	signedData, err := configUpdateEnv.AsSignedData()
	if err != nil {
		return fmt.Errorf("Failed to validate chain creation because config envelope could not be converted to signed data: %s", err)
	}

	err = policy.Evaluate(signedData)
	if err != nil {
		return fmt.Errorf("Failed to validate chain creation, did not satisfy policy: %s", err)
	}

	return nil
}

func (scf *systemChainFilter) inspect(configResources *configResources) error {
	// XXX decide what it is that we will require to be the same in the new config, and what will be allowed to be different
	// Are all keys allowed? etc.

	return nil
}

func (scf *systemChainFilter) authorizeAndInspect(configTx *cb.Envelope) error {
	payload := &cb.Payload{}
	err := proto.Unmarshal(configTx.Payload, payload)
	if err != nil {
		return fmt.Errorf("Rejecting chain proposal: Error unmarshaling envelope payload: %s", err)
	}

	if payload.Header == nil /* || payload.Header.ChannelHeader == nil */ {
		return fmt.Errorf("Rejecting chain proposal: Not a config transaction")
	}

	chdr, err := utils.UnmarshalChannelHeader(payload.Header.ChannelHeader)
	if err != nil {
		return fmt.Errorf("Rejecting chain proposal: Error unmarshaling channel header: %s", err)
	}

	if chdr.Type != int32(cb.HeaderType_CONFIG) {
		return fmt.Errorf("Rejecting chain proposal: Not a config transaction")
	}

	configEnvelope := &cb.ConfigEnvelope{}
	err = proto.Unmarshal(payload.Data, configEnvelope)
	if err != nil {
		return fmt.Errorf("Rejecting chain proposal: Error unmarshalling config envelope from payload: %s", err)
	}

	// Make sure that the config was signed by the appropriate authorized entities
	err = scf.authorize(configEnvelope)
	if err != nil {
		return err
	}

	configResources, err := newConfigResources(configEnvelope)
	if err != nil {
		return fmt.Errorf("Failed to create config manager and handlers: %s", err)
	}

	// Make sure that the config does not modify any of the orderer
	return scf.inspect(configResources)
}
