module github.com/grafana/k6-grpcbin

go 1.21

require (
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/moul/pb v0.0.0-20220425114252-bca18df4138c
	github.com/stretchr/testify v1.8.4
	google.golang.org/grpc v1.59.0
	moul.io/grpcbin v1.0.8
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/glog v1.2.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231127180814-3a041ad873d4 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231127180814-3a041ad873d4 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231127180814-3a041ad873d4 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/grpc-ecosystem/grpc-gateway => github.com/moul/grpc-gateway v1.9.1-0.20190603230725-390f150e109c
