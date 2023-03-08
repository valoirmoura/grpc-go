package main

import (
	pb "github.com/vmcod3r/grpc-go/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, strem pb.CalculatorService_PrimesServer) error {
	log.Printf("primes function was invoked with %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			strem.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
