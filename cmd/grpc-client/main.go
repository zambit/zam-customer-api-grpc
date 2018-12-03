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
	conn, err := grpc.Dial("127.0.0.1:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCustomerAPIServiceGRPCClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.Login(ctx, &pb.LoginRequest{Phone: "+79661861871", Password: "12345"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("Greeting: %#v", r)

}
