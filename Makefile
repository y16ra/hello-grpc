.PHONY: setup
setup:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		hello/hello.proto

.PHONY: clean
clean:
	rm -f hello/hello.pb.go
	rm -f hello/hello_grpc.pb.go
