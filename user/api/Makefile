GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=../srv --micro_out=. --go_out=. proto/user/user.proto

.PHONY: build
build: proto
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o user-api server.go

.PHONY: docker
docker:
	docker build . -t user-api:latest