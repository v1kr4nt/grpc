package main

import (
	"context"
	"log"

	pb "github.com/v1kr4nt/unary/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.ExampleClient

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("error: %s",err)
	}
	client := pb.NewExampleClient(conn)

	req := &pb.HelloRequest{SomeString: "hello from client"}
	client.ServerReply(context.TODO(),req)
}