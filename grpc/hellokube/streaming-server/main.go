package main

import (
	"log"
	"net"
	"time"

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

func (*server) SayHelloStream(req *pb.Request, stream pb.Greeter_SayHelloStreamServer) error {
	for {
		res := &pb.Response{Message: "Hello " + req.GetName()}
		err := stream.Send(res)
		if err != nil {
			log.Printf("could not send response %v: %v", res, err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
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
