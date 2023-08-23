VERSION=0.2.0

build_image:
	docker build -t grafana/k6-grpcbin:${VERSION} .

scan_image:
	docker scan grafana/k6-grpcbin:${VERSION}

push_image:
	docker push grafana/k6-grpcbin:${VERSION}

.PHONY: build_image scan_image
