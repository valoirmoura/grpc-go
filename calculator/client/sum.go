package main

import (
	"context"
	pb "github.com/vmcod3r/grpc-go/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  6,
		SecondNumber: 6,
	})

	if err != nil {
		log.Fatalf("Could not calculator: %v\n", err)
	}

	log.Printf("Calculator: %d\n", res.Result)
}
