package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	pb "github.com/v1kr4nt/unary/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedExampleServer
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("connection error %v", err)
	}
	srv := grpc.NewServer()

	pb.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		log.Fatalf("error: %v", err)
	}
}

func (s *server) ServerReply(c context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("recieve request from client", req.SomeString)
	fmt.Println("hello from server")
	return &pb.HelloResponse{}, errors.New("")
}
