gen-proto:
	protoc -I/usr/local/include -I. \
        -I$(GOPATH)/src \
        --go_out=plugins=grpc:. \
        pb/port/pb.proto

.PHONY: gen-proto