package main

import (
	"log"
	"net"

	pb "github.com/miguellgt/goprojects/grpc/hellokube/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

const port = ":3030"

func (*server) SayHello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println(req)
	return &pb.Response{"Hello " + req.GetName()}, nil
}

func main() {
	log.Println("Helloworld service starting...")
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, new(server))
	log.Println("Helloworld service started successfully")
	s.Serve(ln)
}
