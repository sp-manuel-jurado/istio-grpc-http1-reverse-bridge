ping-grpc:
	grpcurl -v \
-proto ./proto/SP/Rpc/ping_service.proto \
-plaintext \
-import-path ./proto \
-d '{}' \
localhost:10000 \
sp.rpc.PingService/Ping
.PHONY: ping-grpc

ping-http:
	curl -v localhost:10000/ping
.PHONY: ping-http
