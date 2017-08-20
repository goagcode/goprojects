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
	name := flag.String("name", "miguelito", "name about the user to send the server")
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connected %s: %v", address, err)
	}

	client := pb.NewGreeterClient(conn)
	req := &pb.Request{*name}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("could not send request %v: %v", req, err)
	}
	fmt.Println(res)
}
