package client

import (
	"context"
	"time"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/gossip/service"
	pb "github.com/hyperledger/fabric/protos/peer"
	logging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var logger = logging.MustGetLogger("client")
var c pb.SuperviseClient

func startConn() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(3*time.Second))
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(viper.GetString("supervice.address"), opts...)
	if err != nil {
		logger.Infof("did not connect: %v", err)
		return
	}
	c = pb.NewSuperviseClient(conn)
}

// StartSuperviseClient client
func StartSuperviseClient() {
	if c == nil {
		startConn()
	}
	//logger.Infof("%v", channel.GossipChannel.GetPeers())
	gossipService := service.GetGossipService()
	logger.Infof("gossipService:%v", gossipService.Peers())
	peers := gossipService.Peers()
	peerEndpoint, err := peer.GetPeerEndpoint()

	if err != nil {
		//err = fmt.Errorf("Failed to get Peer Endpoint: %s", err)
		logger.Errorf("Failed to get Peer Endpoint: %s", err)
		return
	}

	committer := peer.GetCommitter(util.GetTestChainID())
	var ledgerHeight uint64
	ledgerHeight = 0
	if committer != nil {
		ledgerHeight, err = peer.GetCommitter(util.GetTestChainID()).LedgerHeight()
		if err != nil {
			//err = fmt.Errorf("Failed to get Height: %s", err)
			logger.Errorf("Failed to get Peer Endpoint: %s", err)
		}
		logger.Infof("Height,%v", ledgerHeight)
	}

	logger.Infof("peerEndpoint,%v", peerEndpoint)
	var connpeers []*pb.ConnectPeer
	for _, v := range peers {
		connpeers = append(connpeers, &pb.ConnectPeer{Endpoint: v.Endpoint, Metadata: v.Metadata, PKIid: v.PKIid})
	}
	r, err := c.GetPeer(context.Background(), &pb.PeerInfo{PeerHeight: ledgerHeight, PeerEndpoint: peerEndpoint, ConnectPeers: connpeers})
	if err != nil {
		logger.Infof("could not greet: %v", err)
		return
	}
	logger.Infof("Greeting: %s", r.Message)
}

// RunClient 1s
func RunClient() {
	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for _ = range ticker.C {
			StartSuperviseClient()
		}
	}()
}
