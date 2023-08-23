package main

import (
	"context"
	"testing"

	pb "github.com/moul/pb/grpcbin/go-grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestGRPCServer(t *testing.T) {
	t.Parallel()

	client, connCloseFunc := newGRPCClient(t, ":9000")
	defer connCloseFunc()

	ctx := context.Background()
	res, err := client.DummyUnary(ctx, &pb.DummyMessage{
		FString: "hello",
		FInt32:  42,
	})
	require.NoError(t, err)
	assert.Equal(t, res.FString, "hello")
	assert.Equal(t, res.FInt32, int32(42))
}

type grpcClient struct {
	conn *grpc.ClientConn
	addr string
	pb.GRPCBinClient
}

func newGRPCClient(t *testing.T, addr string) (*grpcClient, func() error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		require.NoError(t, err)
	}
	pbclient := pb.NewGRPCBinClient(conn)
	return &grpcClient{conn: conn, addr: addr, GRPCBinClient: pbclient}, conn.Close
}
