package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/miguellgt/goprojects/grpc/google/google_svc"
	context "golang.org/x/net/context"
	"golang.org/x/net/trace"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGoogleServer(s, new(server))
	s.Serve(lis)
}

type server struct{}

func (*server) Search(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	d := randomDuration(100 * time.Millisecond)
	logSleep(ctx, d)

	select {
	case <-time.After(d):
		return &pb.Result{
			Title: fmt.Sprintf("resutl for [%s] from backend %d", req.Query, *index),
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func logSleep(ctx context.Context, d time.Duration) {
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("sleeping for %s", d)
	}
}

func watch(req *pb.Request, stream pb.Google_WatchServer) error {
	ctx := stream.Context()
	for i := 0; ; i++ {
		d := randomDuration(1 * time.Second)
		logSleep(ctx, d)
		select {
		case <-time.After(d):
			err := stream.Send(&pb.Result{
				Title: fmt.Sprintf("result %d for [%s] from backend %d", i, req.Query, *index),
			})
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
