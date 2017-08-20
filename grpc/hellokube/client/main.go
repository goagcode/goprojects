package main

import (
	"flag"
	"fmt"
	"log"

	pb "github.com/miguellgt/goprojects/grpc/hellokube/Helloworld"
	"golang.org/x/net/context"
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
	req := &pb.Request{*name}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("could not send request %v: %v", req, err)
	}
	fmt.Println(res.GetMessage())
}
