package main

// protoc --go_out=plugins=grpc:chat chat.proto
// protoc --go_out=. chat.proto
import (
	// "example/test_stream_grpc/chat"
	"log"
	"net"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":3692")
	if err != nil {
		log.Fatalf("Failed to listen on port 3692:  %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 3692: %v", err)
	}

}
