package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/miguellgt/goprojects/grpc/google/google_svc"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatalf("could not list to tcp %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGoogleServer(s, new(server))
	s.Serve(lis)
}

type server struct{}

func (*server) Search(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	fmt.Println(req)
	return &pb.Result{
		Title:   "Testing",
		Url:     "googl.com/testing",
		Snippet: "hola",
	}, nil
}
