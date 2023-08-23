package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	grpcbinpb "github.com/moul/pb/grpcbin/go-grpc"
)

// Application serves webpage and grpc reqs without TLS
// Secured connections supported by proxy
var (
	host        = flag.String("host", "localhost", "Host domain name")
	grpcPort    = flag.String("grpc-port", "9000", "The port for grpc server")
	webpagePort = flag.String("webpage-port", "8080", "The port for webpage HTTP server")
)

func main() {
	// parse flags
	flag.Parse()

	// grpc server
	go RunGRPCServerOnAddr(*grpcPort)

	// webpage http server
	methodNames := getMethodNames(grpcbinpb.GRPCBin_serviceDesc.Methods)
	go RunHTTPServerOnAddr(*webpagePort, methodNames)

	// handle Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	log.Fatalf("%s", <-c)
}

func getMethodNames([]grpc.MethodDesc) []string {
	methodNames := []string{}
	for _, method := range grpcbinpb.GRPCBin_serviceDesc.Methods {
		methodNames = append(methodNames, method.MethodName)
	}
	return methodNames
}
