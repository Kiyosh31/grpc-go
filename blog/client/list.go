package main

import (
	"context"
	"grpc-go-course/blog/proto"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c proto.BlogServiceClient) {
	log.Printf("getBlogsList was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}

		log.Println(res)
	}
}
