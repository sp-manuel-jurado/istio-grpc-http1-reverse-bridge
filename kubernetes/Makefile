ping-grpc:
	grpcurl -v \
-proto ../proto/SP/Rpc/ping_service.proto \
-plaintext \
-import-path ../proto \
-d '{}' \
domain.com:80 \
sp.rpc.PingService/Ping
.PHONY: ping-grpc

ping-http:
	curl -v http://domain.com:80/ping
.PHONY: ping-http
