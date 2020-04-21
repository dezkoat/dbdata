package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dezkoat/dbdata/pb"
)

const (
	port = ":50001"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPostServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) CreatePost(
	ctx context.Context,
	in *pb.CreatePostRequest,
) (*pb.CreatePostResponse, error) {
	log.Printf("Adding new post")
	log.Printf("Title: %s", in.GetPost().GetTitle())
	log.Printf("Content: %s", in.GetPost().GetContent())
	return &pb.CreatePostResponse{Post: in.GetPost()}, nil
}

func (s *server) ReadPost(
	ctx context.Context,
	in *pb.ReadPostRequest,
) (*pb.ReadPostResponse, error) {
	log.Printf("Read post")
	log.Printf("id: %s", in.GetId())

	post := &pb.Post{
		Id:               "1",
		Title:            "title",
		Content:          "content",
		TimestampCreated: "kemarin",
		TimestampUpdated: "today",
	}

	return &pb.ReadPostResponse{Post: post}, nil
}

func (s *server) UpdatePost(
	ctx context.Context,
	in *pb.UpdatePostRequest,
) (*pb.UpdatePostResponse, error) {
	log.Printf("Updating post")
	log.Printf("Title: %s", in.GetPost().GetTitle())
	log.Printf("Content: %s", in.GetPost().GetContent())
	return &pb.UpdatePostResponse{Post: in.GetPost()}, nil
}

func (s *server) DeletePost(
	ctx context.Context,
	in *pb.DeletePostRequest,
) (*pb.DeletePostResponse, error) {
	log.Printf("Deleting post")
	log.Printf("id: %s", in.GetId())
	return &pb.DeletePostResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &server{})

	log.Printf("Server started at port %v", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
