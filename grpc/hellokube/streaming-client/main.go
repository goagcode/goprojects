package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	pb "github.com/miguellgt/goprojects/grpc/hellokube/Helloworld"
	"google.golang.org/grpc"
)

const address = "localhost:3030"

func main() {
	server := flag.String("server", "127.0.0.1:3030", "Server address.")
	name := flag.String("name", "miguellgt", "username to use.")

	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connected %s: %v", *server, err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	stream, err := client.SayHelloStream(context.Background(), &pb.Request{Name: *name})
	if err != nil {
		log.Fatalf("could not get streaming: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not get message: %v", err)
		}
		fmt.Println(response.GetMessage())
	}
}
