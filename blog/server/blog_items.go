package main

import (
	"grpc-go-course/blog/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"authorId"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlog(data *BlogItem) *proto.Blog {
	return &proto.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
