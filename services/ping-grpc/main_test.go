package main_test

import (
	"context"
	"testing"

	pb "github.com/socialpoint/istio-grpc-http1-reverse-bridge/services/ping-grpc/pkg/sp_rpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestServiceServer_Ping(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	conn, err := grpc.Dial("envoy:10001", grpc.WithInsecure())
	r.NoError(err)

	cli := pb.NewPingServiceClient(conn)
	_, err = cli.Ping(context.Background(), &empty.Empty{})
	a.NoError(err)
}
