package main

import (
	"context"
	"fmt"
	"grpc-go-course/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *proto.Blog) (*proto.BlogId, error) {
	log.Printf("CreateBlog was invoekd with: %v", in)

	data := BlogItem{
		AuthorId: in.GetAuthorId(),
		Title:    in.GetTitle(),
		Content:  in.GetContent(),
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &proto.BlogId{
		Id: oid.Hex(),
	}, nil
}
