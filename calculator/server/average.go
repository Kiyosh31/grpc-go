package main

import (
	"grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Average(stream proto.CalculatorService_AverageServer) error {
	log.Printf("Average was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.AverageResponse{
				Result: int32(sum) / int32(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %s", err)
		}

		log.Printf("Receiving number: %d \n", req.GetNumber())

		sum = req.GetNumber()
		count++
	}
}
