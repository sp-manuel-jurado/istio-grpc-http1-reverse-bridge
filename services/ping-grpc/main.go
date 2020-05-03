package main

import (
	"context"
	"log"
	"net"

	pb "github.com/socialpoint/istio-grpc-http1-reverse-bridge/services/ping-grpc/pkg/sp_rpc"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	addr = ":10005"
)

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPingServiceServer(grpcServer, serviceServer{})
	log.Printf("grpc server in localhost" + addr + " started")

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}

type serviceServer struct {
}

func (s serviceServer) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
