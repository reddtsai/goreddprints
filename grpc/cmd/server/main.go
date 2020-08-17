package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/reddtsai/goreddprints/grpc/pkg/rpc"
)

func main() {
	lis, err := net.Listen("tcp", ":6424")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	fmt.Println("grpc sample service ...")
	rpc.RegisterSampleServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}

type server struct {
	rpc.SampleServiceServer
}

func (s *server) SayHello(ctx context.Context, in *rpc.HelloRequest) (*rpc.HelloReply, error) {
	return &rpc.HelloReply{Message: "Hello " + in.GetName()}, nil
}
