package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/reddtsai/goreddprints/grpc/pkg/rpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:6424", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := rpc.NewSampleServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &rpc.HelloRequest{Name: "redd"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.GetMessage())
}
