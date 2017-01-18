package client

import (
	"context"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/gossip/service"
	pb "github.com/hyperledger/fabric/protos/peer"
	logging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var logger = logging.MustGetLogger("client")

// StartSuperviseClient client
func StartSuperviseClient() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(10*time.Second))
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(viper.GetString("supervice.address"), opts...)
	if err != nil {
		logger.Infof("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	logger.Infof("%v", c)
	// Contact the server and print out its response.
	name := defaultName
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})

	//logger.Infof("%v", channel.GossipChannel.GetPeers())
	gossipService := service.GetGossipService()
	logger.Infof("gossipService:%v", gossipService.Peers())
	if err != nil {
		logger.Infof("could not greet: %v", err)
		return
	}
	peerEndpoint, err := peer.GetPeerEndpoint()

	if err != nil {
		err = fmt.Errorf("Failed to get Peer Endpoint: %s", err)
	}

	committer := peer.GetCommitter(util.GetTestChainID())
	if committer != nil {
		ledgerHeight, err := peer.GetCommitter(util.GetTestChainID()).LedgerHeight()
		if err != nil {
			err = fmt.Errorf("Failed to get Height: %s", err)
		}
		logger.Infof("Height,%v", ledgerHeight)
	}

	logger.Infof("peerEndpoint,%v", peerEndpoint)

	logger.Infof("Greeting: %s", r.Message)
}

// RunClient 1s
func RunClient() {
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for _ = range ticker.C {
			StartSuperviseClient()
		}
	}()
}
