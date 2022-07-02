package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-example/example"
	"log"
	"net"
	"time"
)

func main() {
	startClient()
	//startServer()
}

func startServer() {
	// Start net listener for grpc server
	log.Println("Starting listening on port 8080")
	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)

	gsrv := grpc.NewServer()

	// Implement Server Stub from generated files
	var esrv example.Testserver
	example.RegisterEchoServer(gsrv, esrv)

	if err := gsrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startClient() {
	// open connection to server
	conn, err := grpc.Dial("grpc-server:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("lets go")
	// instantiate client from generated stubs
	client := example.NewEchoClient(conn)
	ctx := context.Background()
	test := example.Test{
		Id:          12,
		Time:        timestamppb.Now(),
		Description: "sfkjßdklfkß   ßdlfjklßd",
	}
	for true {
		time.Sleep(2 * time.Second)
		timestamp, err := client.GetTimestamp(ctx, &example.Empty{})
		if err != nil {
			return
		}
		fmt.Println("aktueller timestamp: ", timestamp.GetTime().String())
		time.Sleep(1 * time.Second)
		hello, err := client.Hello(ctx, &test)
		if err != nil {
			return
		}
		fmt.Println("client: ", hello)
	}
}
