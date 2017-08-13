package main

import (
	"fmt"
	"log"

	pb "github.com/miguellgt/goprojects/grpc/helloworld/helloworld_service"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetingClient(conn)
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "Miguel Angel Galicia",
		Age:  29,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
