// #### add by chenqiao
package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	pb "github.com/hyperledger/fabric/protos/peer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct{}

// func init() {
// 	go StartServer()
// }

type ClientInfo struct {
	peerIp string
	name   string
}

func (s *server) ProcessMessage(ctx context.Context, inputMessage *pb.MessageInput) (*pb.MessageOutput, error) {

	//message
	peerMessage := NewPeerMessage()
	peerMessage.InitMessage(inputMessage)

	PeerStatusMap[peerMessage.PeerIp] = peerMessage

	fmt.Println("@@@@chenqiao: ", PeerStatusMap)

	return &pb.MessageOutput{Output: "hehe " + "I GOT IT"}, nil
}

func QueryClient(blockindex uint64) *pb.Mblock {
	address := beego.AppConfig.String("queryServer")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		// log.Panic("can't connect to server %v", err)
		fmt.Printf("@@@@@@ chenqiao: cannot dial %s , The error is %v\n", address, err)

	}
	c := pb.NewQueryPeerClient(conn)
	r, err := c.QueryMessage(context.Background(), &pb.QueryBlocks{BlockIndex: blockindex})
	if err != nil {
		// log.Panic("could not greet %v", err)
		fmt.Println("@@@@@@ chenqiao: cannot execute processMessage, the error ", err)
	}
	// log.Printf("Greeting: %s", r.Output)
	// fmt.Println("@@@@@@ chenqiao: Greeeting: ", r)
	return r
}

// const (
// 	port = ":38254"
// )

func StartServer() {
	port := beego.AppConfig.String("rpcServerPost")
	lis, err := net.Listen("tcp", port)
	// fmt.Println("i am come here!!!")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStatusPeerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	// reflection.RegisterStatusPeerServer(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		// fmt.Println("There is error %v", err)
		// continue
	}
}
