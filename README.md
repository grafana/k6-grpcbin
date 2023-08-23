# k6 gRPCBin Play: gRPC Request & Response Service

A simple gRPC service to experiment with the k6 gRCP API, deployed at https://grpcbin.test.k6.io/

This project is based on [grpcbin](https://github.com/moul/grpcbin) by [Manfred Touron](https://v1.manfred.life/).

## Run in Docker

You can deploy or run this project locally using the [`grafana/k6-grpcbin` Docker Image](https://hub.docker.com/r/grafana/k6-grpcbin):

```bash
docker pull grafana/k6-grpcbin
docker run -it --rm -p 9000:9000 -p 8080:8080 grafana/k6-grpcbin
```

or building the Docker image:

```bash
docker build -t k6-grpcbin-local .
docker run -it --rm -p 9000:9000 -p 8080:8080 k6-grpcbin-local
```

The project runs an insecure gRPC service, without TLS.
