package main

import (
	"context"
	"log"
	"os"
	"time"

	pbFile "github.com/farhadmirr/proto"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:9090"
	defaultName     = "Farhad"
	defaultLastname = "Mir"
)

func main() {
	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Can not dial to :%v", err)
	}
	defer connection.Close()
	client := pbFile.NewMyServerClient(connection)
	clientName := defaultName
	clientLname := defaultLastname
	if len(os.Args) > 1 {
		clientName = os.Args[1]
		clientLname = os.Args[2]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	serverResponse, err := client.SayHello(ctx, &pbFile.HelloBody{Name: clientName, Lastname: clientLname})
	if err != nil {
		log.Fatalf("Unable to say hello | %v", err)
	}
	log.Printf("Server Response: %s", serverResponse.GetRespond())

}
