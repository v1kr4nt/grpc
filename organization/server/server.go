package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/v1kr4nt/organization/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedInitOrgServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("error %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterInitOrgServiceServer(srv, &server{})

	if e := srv.Serve(lis); e != nil {
		log.Fatalf("error: %v", e)
	}
}

func (s *server) InitOrganization(ctx context.Context, req *pb.InitOrgRequest) (*pb.InitOrgResponse, error) {
	log.Println(req)
	return &pb.InitOrgResponse{
		Id: req.Id,
	}, errors.New("")
}
