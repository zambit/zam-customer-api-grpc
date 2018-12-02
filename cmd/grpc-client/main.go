package main

import (
	"context"
	"log"
	"time"

	"git.zam.io/microservices/customer-api/pb"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCustomerAPIServiceGRPCClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Login(ctx, &pb.LoginRequest{Id: 5})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %#v", r)

}
