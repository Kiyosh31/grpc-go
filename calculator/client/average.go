package main

import (
	"context"
	"grpc-go-course/calculator/proto"
	"log"
)

func doAverage(c proto.CalculatorServiceClient) {
	log.Printf("doAverage was invoked")

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while stream: %s", err)
	}

	numbers := []int32{3, 5, 8, 15, 20}

	for _, number := range numbers {
		log.Printf("Sending number: %d \n", number)

		stream.Send(&proto.AverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving res: %s", err)
	}

	log.Printf("Average: %f \n", res.GetResult())
}
