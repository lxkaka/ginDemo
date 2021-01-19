package main

import (
	"fmt"
	"ginDemo/api"
	"io"
	"net"

	"google.golang.org/grpc"
)

type TaskServer struct {
	api.UnimplementedTaskHandlerServer
}

func (s *TaskServer) SubmitTask(stream api.TaskHandler_SubmitTaskServer) (err error) {
	for {
		in, e := stream.Recv()
		if e == io.EOF {
			return nil
		}
		if e != nil {
			return e
		}
		fmt.Printf("server receive: %s", in.Data)
		out := fmt.Sprintf("res-%s", in.Data)
		if e = stream.Send(&api.TaskResponse{Identifier: in.Identifier, Sequence: in.Sequence, Data: []byte(out)}); e != nil {
			return e
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		panic(err)
	}
	gprcServer := grpc.NewServer()
	api.RegisterTaskHandlerServer(gprcServer, &TaskServer{})
	gprcServer.Serve(lis)
}
