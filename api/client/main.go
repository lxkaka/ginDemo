package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ginDemo/api"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewMyServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.User(ctx, &api.UserReq{UserIDs: []int32{1, 2}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("gprc result: %+v", r.Data)
}
