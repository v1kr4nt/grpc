package main

import (
	"context"
	"log"

	pb "github.com/v1kr4nt/organization/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	client := pb.NewInitOrgServiceClient(conn)

	req := &pb.InitOrgRequest{
		Id:      "1234",
		OrgName: "test",
	}
	client.InitOrganization(context.TODO(), req)
}
