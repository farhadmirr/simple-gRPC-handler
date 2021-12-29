//this is server side file
//to run client side run mainClient.go

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pbFile "github.com/farhadmirr/proto"
	"google.golang.org/grpc"
)

const port = ":9090"

type server struct {
	pbFile.UnimplementedMyServerServer
}

func (s *server) SayHello(ctx context.Context, income *pbFile.HelloBody) (*pbFile.HelloReply, error) {
	fmt.Println("Received: %v", income.GetName(), income.GetLastname())
	return &pbFile.HelloReply{Respond: "Hello  " + income.GetName() + " " + income.GetLastname()}, nil
}
func main() {
	fmt.Println("Preparing to listen on port" + port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Can not listen on port " + port)
	}
	grpc := grpc.NewServer()
	pbFile.RegisterMyServerServer(grpc, &server{})
	fmt.Println("Server is ready to listen at port" + port)
	if err := grpc.Serve(listener); err != nil {
		log.Fatalf("Failed to serve")
	}
}
