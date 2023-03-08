package main

import (
	"context"
	pb "github.com/vmcod3r/grpc-go/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet as invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Valoir",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
