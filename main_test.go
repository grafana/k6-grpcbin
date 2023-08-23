package main

import (
	"log"
	"net"
	"os"
	"testing"
	"time"

	grpcbinpb "github.com/moul/pb/grpcbin/go-grpc"
)

func TestMain(m *testing.M) {
	if !startService() {
		log.Println("Timed out waiting for HTTP to come up")
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func startService() bool {
	webpagePort := "8080"
	grpcPort := "9000"
	methodNames := getMethodNames(grpcbinpb.GRPCBin_serviceDesc.Methods)

	go RunHTTPServerOnAddr(webpagePort, methodNames)

	ok := waitForPort(":" + webpagePort)
	if !ok {
		log.Println("Timed out waiting for HTTP server to come up")
		return false
	}

	go RunGRPCServerOnAddr(grpcPort)

	ok = waitForPort(":" + grpcPort)
	if !ok {
		log.Println("Timed out waiting for GRPC server to come up")
		return false
	}

	return true
}

func waitForPort(address string) bool {
	waitChan := make(chan struct{})

	go func() {
		for {
			conn, err := net.DialTimeout("tcp", address, time.Second)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			if conn != nil {
				waitChan <- struct{}{}
				return
			}
		}
	}()

	timeout := time.After(5 * time.Second)
	select {
	case <-waitChan:
		return true
	case <-timeout:
		return false
	}
}
