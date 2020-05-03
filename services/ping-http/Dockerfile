FROM golang:1.13

WORKDIR /go/github.com/socialpoint/istio-grpc-http1-reverse-bridge/services/ping-grpc
COPY . .

EXPOSE 8888

CMD ["go", "run", "main.go"]
