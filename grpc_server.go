package main

import (
	"fmt"
	"log"
	"net"

	abepb "github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	abehandler "github.com/grpc-ecosystem/grpc-gateway/examples/server"
	addsvcpb "github.com/moul/pb/addsvc/go-grpc"
	grpcbinpb "github.com/moul/pb/grpcbin/go-grpc"
	hellopb "github.com/moul/pb/hello/go-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	addsvchandler "moul.io/grpcbin/handler/addsvc"
	grpcbinhandler "moul.io/grpcbin/handler/grpcbin"
	hellohandler "moul.io/grpcbin/handler/hello"
)

func RunGRPCServerOnAddr(grpcPort string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failted to listen: %v", err)
	}

	// create gRPC server
	s := grpc.NewServer()
	grpcbinpb.RegisterGRPCBinServer(s, &grpcbinhandler.Handler{})
	hellopb.RegisterHelloServiceServer(s, &hellohandler.Handler{})
	addsvcpb.RegisterAddServer(s, &addsvchandler.Handler{})
	abepb.RegisterABitOfEverythingServiceServer(s, abehandler.NewHandler())
	// register reflection service on gRPC server
	reflection.Register(s)

	// serve
	log.Printf("listening on %s (insecure gRPC)\n", grpcPort)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
