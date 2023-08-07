package main

import (
	"context"
	"fmt"
	"grpc-go-course/blog/proto"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogs(in *empty.Empty, stream proto.BlogService_ListBlogsServer) error {
	log.Printf("ListBlogs was invoked with: %v", in)

	curr, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		data := &BlogItem{}
		err := curr.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from mongo: %v", err),
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = curr.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal: %v", err),
		)
	}

	return nil
}
