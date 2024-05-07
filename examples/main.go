package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lattice-manager-grpc/gen/helloworld"
	tblockv1 "lattice-manager-grpc/gen/tblock/v1"
	"log"
	"time"
)

func main() {
	// initial gRPC client
	conn, fn := newClient()
	defer fn(conn)

	SayHello(conn)
	tBlockDetails(conn)
}

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func newClient() (*grpc.ClientConn, func(conn *grpc.ClientConn)) {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn, func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}
}

func SayHello(conn *grpc.ClientConn) {
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "Hello"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func tBlockDetails(conn *grpc.ClientConn) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tBlockClient := tblockv1.NewTBlockServiceClient(conn)
	r, err := tBlockClient.Details(ctx, &tblockv1.TBlockServiceDetailsRequest{
		Hash: "0xeKWUdQIZPMlb2ApNjByhREJm8fgcBDHvt790wLTiOCqa5GFSxo1VXrzYn6s3koa4",
	})
	if err != nil {
		log.Fatalf("Could not get block details: %v", err)
	} else {
		log.Printf("Block details: %v", r)
	}
}
