package node

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwset"
	"github.com/hyperledger/fabric/core/ledger/ledgermgmt"
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type server struct{}

type argsData struct {
	Args []string `json:"args"`
}

func getChainCodeID(inputs string) string {
	s := strings.Split(inputs, "\n")
	cn := strings.Split(s[1], ":")

	chainCodeID := string(cn[0][1:])

	return chainCodeID

}

func getInputDetail(input [][]byte) string {
	// func chainname path funcArgs
	var ans string

	//deploy

	if string(input[0]) == "deploy" {
		ans = ans + string(input[0]) + " " + string(input[1]) + " "

		s := strings.Split(string(input[2]), "\n")
		for i := 0; i < len(s); i++ {
			ans = ans + s[i] + " "
		}
	} else {
		for i := 0; i < len(input); i++ {
			ans = ans + string(input[i]) + " "
		}
	}

	//invoke

	return ans
}

func getBlockDataByIndex(j int, blocks []*common.Block) *pb.Mblock {
	blocksheader := blocks[j].GetHeader()
	blocksData := blocks[j].GetData()

	blockMetadata := blocks[j].GetMetadata()
	nowHash := base64.StdEncoding.EncodeToString(blocksheader.Hash())
	preHash := base64.StdEncoding.EncodeToString(blocksheader.PreviousHash)
	dataHash := base64.StdEncoding.EncodeToString(blocksheader.DataHash)

	mheader := &pb.MblockHeader{
		Number:       blocksheader.Number,
		PreviousHash: preHash,
		DataHash:     dataHash,
		NowHash:      nowHash,
	}

	tmpblockData := []*pb.TransData{}

	for k := 0; k < len(blocksData.Data); k++ {
		//尝试解析Data得到txid和chainid
		e, err := utils.GetEnvelopeFromBlock(blocksData.Data[k])
		if err != nil {
			fmt.Println("@@@@chenqiao: Envelope err", err)
			continue
		}
		p, err := utils.UnmarshalPayload(e.Payload)
		if err != nil {
			fmt.Println("@@@@chenqiao: Payload err", err)
			continue
		}

		hdr, err := utils.UnmarshalChannelHeader(p.Header.ChannelHeader)
		if err != nil {
			fmt.Println("@@@@chenqiao: ChannelHeader err", err)
			continue
		}

		txid := hdr.TxId
		tchainID := hdr.ChannelId
		time := hdr.Timestamp

		// fmt.Println("@@@@ chenqiao txid: ", txid)
		// fmt.Println("@@@@ chenqiao tchainID: ", tchainID)
		// fmt.Println("@@@@ chenqiao Time: ", time)

		trans := &pb.Transaction{}
		transerr := proto.Unmarshal(p.Data, trans)
		if transerr != nil {
			fmt.Println("ERROR!!!!!")
			transData := &pb.TransData{
				Txid:    txid,
				ChainID: tchainID,
				Time:    time,
			}
			tmpblockData = append(tmpblockData, transData)
			continue
		}

		// fmt.Println("@@@@ chenqiao transVersion", trans.Version)
		// fmt.Println("@@@@ chenqiao Time!", trans.Timestamp)

		cap := &pb.ChaincodeActionPayload{}
		caperr := proto.Unmarshal(trans.Actions[0].Payload, cap)
		if caperr != nil {
			fmt.Println("CAP ERROR!!!!!")
			continue
		}

		cpp := &pb.ChaincodeProposalPayload{}
		cpperr := proto.Unmarshal(cap.ChaincodeProposalPayload, cpp)
		if cpperr != nil {
			fmt.Println("CPP ERROR!!!!!")
			continue
		}

		prpayload := &pb.ProposalResponsePayload{}
		prperr := proto.Unmarshal(cap.Action.ProposalResponsePayload, prpayload)
		if prperr != nil {
			fmt.Println("PRP ERROR!!! ")
			transData := &pb.TransData{
				Txid:    txid,
				ChainID: tchainID,
				Time:    time,
			}
			tmpblockData = append(tmpblockData, transData)
			continue
		}

		ccids := &pb.ChaincodeAction{}
		ccidserr := proto.Unmarshal(prpayload.Extension, ccids)
		if ccidserr != nil {
			fmt.Println("CCIDS ERROR!!! ")
			continue
		}

		cis := &pb.ChaincodeInvocationSpec{}
		ciserr := proto.Unmarshal(cpp.Input, cis)
		if ciserr != nil {
			fmt.Println("CIS ERROR!!!!!")
			continue
		}

		inputDetail := getInputDetail(cis.ChaincodeSpec.Input.Args)

		// fmt.Println("UUUUUUUUUUUUU", cis.ChaincodeSpec.Input)
		// a, err := json.Marshal(cis.ChaincodeSpec.Input)
		// if err != nil {
		// 	fmt.Println("json err!!! ", err)
		// }

		// fmt.Println("hehehehelalalala", string(a))

		// var tdata argsData
		// if err := json.Unmarshal(a, &tdata); err == nil {
		// 	for i := 0; i < len(tdata.args); i++ {
		// 		abytes, err64 := base64.StdEncoding.DecodeString(tdata.args[i])
		// 		if err64 != nil {
		// 			fmt.Println("base64 err", err)
		// 		}
		// 		fmt.Println("HOHOHOHOHOHOHO", string(abytes))
		// 	}

		// }

		txRWSet := &rwset.TxReadWriteSet{}
		txRWSet.Unmarshal(ccids.Results)

		// fmt.Println("HEHEHEHEHEHEHEHE!!!HE!!", inputDetail)

		chainCodeID := getChainCodeID(txRWSet.String())

		// ccids.Results
		var result string
		for _, nsRWSet := range txRWSet.NsRWs {
			if nsRWSet.NameSpace == chainCodeID {
				for _, value := range nsRWSet.Writes {
					// fmt.Println("LLLLL ", value.Key, "QQQQQ", string(value.Value))
					result = result + value.Key + ":" + string(value.Value) + "  "
				}
			}

		}

		che := &pb.ChaincodeHeaderExtension{}
		cheerr := proto.Unmarshal(hdr.Extension, che)
		if cheerr != nil {
			fmt.Println("CHE ERROR!!! ")
			continue
		}

		ms := &msp.SerializedIdentity{}

		hsr, hsrerr := utils.GetSignatureHeader(p.Header.SignatureHeader)
		if hsrerr != nil {
			fmt.Println("HSR ERROR!!! ")
			continue
		}

		mserr := proto.Unmarshal(hsr.Creator, ms)
		if mserr != nil {
			fmt.Println("MS ERROR!!! ")
			continue
		}

		transData := &pb.TransData{
			Txid:        txid,
			ChainID:     tchainID,
			Time:        time,
			ChainCodeID: chainCodeID,
			Result:      result,
			Input:       inputDetail,
			// Payload:     chainpayload,
			Type:      strconv.Itoa(int(hdr.Type)),
			Nonce:     base64.StdEncoding.EncodeToString(hsr.Nonce),
			Signature: base64.StdEncoding.EncodeToString(e.Signature),
		}
		tmpblockData = append(tmpblockData, transData)

		// tmpblockData = append(tmpblockData, base64.StdEncoding.EncodeToString(blocksData.Data[k]))
		// fmt.Println("@@@@chenqiao blocksData: ", base64.StdEncoding.EncodeToString(blocksData.Data[k]))
	}

	tmpblockMetaData := []string{}

	for k := 0; k < len(blockMetadata.Metadata); k++ {

		tmpblockMetaData = append(tmpblockMetaData, base64.StdEncoding.EncodeToString(blockMetadata.Metadata[k]))
		// fmt.Println("@@@@chenqiao blocksMetaData: ", base64.StdEncoding.EncodeToString(blockMetadata.Metadata[k]))
	}

	return &pb.Mblock{
		Header: mheader,
		Data: &pb.MblockData{
			Datas: tmpblockData,
		},
		Metadata: &pb.MblockMetadata{
			Metadata: tmpblockMetaData,
		},
	}
}

func (s *server) QueryMessage(ctx context.Context, query *pb.QueryBlocks) (*pb.Mblock, error) {

	blockindex := query.BlockIndex
	blockids := []uint64{blockindex}

	chainID := util.GetTestChainID()
	commit := peer.GetCommitter(chainID)

	blocks := commit.GetBlocks(blockids)

	blockinfo := getBlockDataByIndex(0, blocks)

	return blockinfo, nil
}

func PeerServer() {
	port := viper.GetString("peer.statusPeer.grpcServerPort") //os.Getenv("CORE_PEER_GRPCPORTS")
	lis, err := net.Listen("tcp", port)
	// fmt.Println("i am come here!!!")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterQueryPeerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	// reflection.RegisterStatusPeerServer(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		// fmt.Println("There is error %v", err)
		// continue
	}
}

func StatusClient() {
	// fmt.Println("@@@@@chenqiao test:  ", viper.GetString("peer.networkId"))
	address := viper.GetString("peer.statusPeer.sendGrpcServer")       //os.Getenv("CORE_PEER_GRPCSERVER")
	strTimeCycle := viper.GetString("peer.statusPeer.sendStatusCycle") //os.Getenv("CORE_PEER_SENDGRPC_TIME")
	timeCycle, err := strconv.Atoi(strTimeCycle)
	if err != nil {
		panic(err)
	}
	// fmt.Println("@@@@chenqiao peer: ", timeCycle)
	// cid := "**TEST_CHAINID**"

	// defer conn.Close()
	for {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		// fmt.Println("@@@@chenqiao peer: ", timeCycle)
		if err != nil {
			// log.Panic("can't connect to server %v", err)
			fmt.Printf("@@@@@@ chenqiao: cannot dial %s , The error is %v\n", address, err)
			conn.Close()
			time.Sleep(time.Duration(timeCycle) * 1e9)

			// conn, err = grpc.Dial(address, grpc.WithInsecure())
			continue
		}

		input := getPeerStatus()

		c := pb.NewStatusPeerClient(conn)
		r, err := c.ProcessMessage(context.Background(), input)
		if err != nil {
			// log.Panic("could not greet %v", err)
			fmt.Println("@@@@@@ chenqiao: cannot execute processMessage, the error ", err)
			conn.Close()
			time.Sleep(time.Duration(timeCycle) * 1e9)
			// conn, err = grpc.Dial(address, grpc.WithInsecure())
			continue
		}
		// log.Printf("Greeting: %s", r.Output)
		fmt.Println("@@@@@@ chenqiao: Greeeting: ", r.Output)
		time.Sleep(time.Duration(timeCycle) * 1e9)
		conn.Close()
	}
}

func getPeerStatus() *pb.ChannelMessage {
	ledgerIds, err := ledgermgmt.GetLedgerIDs()
	if err != nil {
		return nil
	}

	cM := &pb.ChannelMessage{}

	for _, chainID := range ledgerIds {
		fmt.Println("@@@@@@@chenqiao: chainID ", chainID)
		commit := peer.GetCommitter(chainID)
		if commit == nil {
			fmt.Println("@@@@@chenqiao: the commit is nil, don't care")
			continue
		}

		//获取当前节点的peer身份
		// peerId := peer.GetLocalIP()

		peerName := os.Getenv("CORE_PEER_ID")
		peerId := os.Getenv("CORE_PEER_ADDRESS")
		// peerId := peerName + peerKey

		//链信息
		height, err := commit.LedgerHeight()
		if err != nil {
			fmt.Println("@@@@@@@ chenqiao: no height !!!!!!")
			continue
		}

		//块信息
		blocksids := []uint64{}
		var i uint64

		//只获取最新的100块
		if height > 100 {
			i = height - 100
		} else {
			i = 0
		}
		for ; i < height; i++ {
			blocksids = append(blocksids, uint64(i))
		}

		//封装消息
		ans := &pb.MessageInput{}

		ans.PeerIp = peerId
		ans.Height = height
		ans.PeerName = peerName
		ans.ChannelID = chainID

		mblock := []*pb.Mblock{}

		if height != 0 {

			blocks := commit.GetBlocks(blocksids)

			for j := 0; j < len(blocks); j++ {
				tmpmblock := getBlockDataByIndex(j, blocks)
				mblock = append(mblock, tmpmblock)

			}

		}

		ans.Mblocks = mblock
		cM.ChannelInput = append(cM.ChannelInput, ans)

	}

	return cM
}
