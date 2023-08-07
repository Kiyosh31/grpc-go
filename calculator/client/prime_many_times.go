package main

import (
	"context"
	"grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrimeManyTimes(c proto.CalculatorServiceClient) {
	log.Printf("doPrimeManyTimes was invoked")

	req := proto.PrimeRequest{
		Num1: 123456789,
	}

	stream, err := c.Primes(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while primes: %s", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %s", err)
		}

		log.Printf("Primes: %d\n", res.GetResult())
	}
}
