package main

import (
	"context"
	"fmt"
	"ginDemo/api"
	"io"
	"log"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:5001"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewTaskHandlerClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	st, err := c.SubmitTask(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err := st.Recv()
			if err == io.EOF {
				fmt.Printf("server close")
				close(waitc)
			}
			if err != nil {
				log.Fatalf("receive err:%v", err)
			}
			if res != nil {
				fmt.Printf("gprc result: %+v", string(res.Data))
			}
		}
	}()

	var (
		input string
		count int32
	)
	for {
		count += 1
		fmt.Scan(&input)
		err = st.Send(&api.TaskRequest{Identifier: "test", Sequence: count, Data: []byte(input)})
		if err != nil {
			log.Fatal(err)
		}
		if count > 5 {
			break
		}
	}
	st.CloseSend()
	<-waitc
}
