package main

import (
	"context"
	"grpc-go-course/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *proto.BlogId) (*proto.Blog, error) {
	log.Printf("ReadBlog was invoked with: %v", in)

	oid, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := BlogItem{}
	filter := bson.M{
		"_id": oid,
	}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(&data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with id provided",
		)
	}

	return documentToBlog(&data), nil
}
