package main

import (
	"context"
	"grpc-go-course/blog/proto"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	addr       = "0.0.0.0:50051"
	collection *mongo.Collection
)

type Server struct {
	proto.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Could not connect to mongodb: %s", err)
	}
	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("Could not connect to mongodb: %s", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on: %s\n", addr)

	svc := grpc.NewServer()
	proto.RegisterBlogServiceServer(svc, &Server{})

	if err = svc.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
