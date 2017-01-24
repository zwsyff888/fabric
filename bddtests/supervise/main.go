package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/hyperledger/fabric/protos/peer"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
)

const (
	port = ":6060"
)

var logger = logging.MustGetLogger("client")

var Peers map[string]*Peer

var PeersSlice []*Peer

// server is used to implement helloworld.GreeterServer.
type server struct{}

type Peer struct {
	Peername  string       `json:"peername"`
	Peerinfo  *pb.PeerInfo `json:"peerinfo"`
	Timestamp int64        `json:"timestamp"`
	Alive     bool         `json:"alive"`
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetPeer(ctx context.Context, in *pb.PeerInfo) (*pb.PeerReply, error) {
	//logger.Infof("PeerEndpoint,%v", in)
	key := "NAME:" + in.PeerEndpoint.ID.Name + " ADDRESS:" + in.PeerEndpoint.Address
	_, ok := Peers[key]
	if ok {
		Peers[key].Peername = in.PeerEndpoint.ID.Name
		Peers[key].Peerinfo = in
		Peers[key].Timestamp = time.Now().Unix()
		Peers[key].Alive = true
	} else {
		Peers[key] = &Peer{Peername: in.PeerEndpoint.ID.Name, Peerinfo: in, Timestamp: time.Now().Unix(), Alive: true}
		PeersSlice = append(PeersSlice, Peers[key])
	}
	logger.Infof("Peers:%v,%v", Peers[key].Timestamp, Peers[key].Peername)
	return &pb.PeerReply{Message: "GetPeer Success " + in.PeerEndpoint.ID.Name}, nil
}

// Echo websocket
func Echo(ws *websocket.Conn) {
	logger.Infof("websocket Echo")
	var err error

	for {

		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		if reply == "secret id" {
			json, _ := json.Marshal(PeersSlice)
			logger.Infof("Peers0:%v,%v", PeersSlice[0].Timestamp, PeersSlice[0].Peername)
			if err := websocket.Message.Send(ws, string(json)); err != nil {
				fmt.Println("Can't send")
				break
			}

		}
	}
}

func main() {
	Peers = make(map[string]*Peer)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSuperviseServer(grpcServer, &server{})
	// Register reflection service on gRPC server.
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for _ = range ticker.C {
			ctime := time.Now().Unix()
			for i := range Peers {
				if ctime-Peers[i].Timestamp > 3 {
					Peers[i].Alive = false
				}
			}
		}
	}()
	fmt.Println("begin")
	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line
	http.Handle("/socket", websocket.Handler(Echo))
	go func() {

		if err := http.ListenAndServe(":7056", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}

		fmt.Println("end")
	}()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
