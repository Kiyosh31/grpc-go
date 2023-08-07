package main

import (
	"grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Primes(in *proto.PrimeRequest, stream proto.CalculatorService_PrimesServer) error {
	log.Printf("Primes was invoekd with %s", in)

	number := in.GetNum1()
	divisor := int64(2)

	for number > 1 {
		if number%int32(divisor) == 0 {
			stream.Send(&proto.PrimeResponse{
				Result: int32(divisor),
			})

			number /= int32(divisor)
		} else {
			divisor++
		}
	}

	return nil
}
