package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/miguellgt/goprojects/grpc/hellogrpc/messages"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

type server struct{}

const port = ":3030"

func (*server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println(req)
	return &pb.HelloResponse{
		Message: "Hello Bro: " + req.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, new(server))
	s.Serve(lis)
}
