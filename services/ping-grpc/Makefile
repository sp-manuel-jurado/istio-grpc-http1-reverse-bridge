up:
	docker-compose up
.PHONY: up

test:
	go clean -testcache && go test -v ./...
.PHONY: test

docker-login:
	docker build -t spmanueljurado/ping-grpc:latest .
.PHONY: build

docker-build:
	docker build -t spmanueljurado/ping-grpc:latest .
.PHONY: build

docker-push:
	 docker push spmanueljurado/ping-grpc:latest
.PHONY: docker-push
