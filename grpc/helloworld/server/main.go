package main

import (
	"fmt"
	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/miguellgt/goprojects/grpc/helloworld/helloworld_service"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println(in)
	return &pb.HelloResponse{
		Greeting: "Hello " + in.GetName() + " have a nice day",
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8080")
	s := grpc.NewServer()
	pb.RegisterGreetingServer(s, &server{})
	s.Serve(lis)
}
