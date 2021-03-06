package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/miguellgt/goprojects/grpc/google/google_svc"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

func main() {
	server := flag.String("s", "localhost:3030", "address to grpc server")
	query := flag.String("q", "gRPC framework", "query to search int the google search")
	mode := flag.String("m", "search", "mode to use the service")

	flag.Parse()

	// Connect to the server
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to the server %s: %v", *server, err)
	}
	defer conn.Close()

	client := pb.NewGoogleClient(conn)

	// Run the RPC.
	switch *mode {
	case "search":
		search(client, *query)
	case "watch":
		watch(client, *query)
	default:
		log.Fatalf("Unknown mode: %q", *mode)
	}
}

func search(client pb.GoogleClient, query string) {
	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancel()

	req := &pb.Request{Query: query}
	res, err := client.Search(ctx, req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func watch(client pb.GoogleClient, query string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &pb.Request{Query: query}
	stream, err := client.Watch(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("and now your watch is ended")
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
}
