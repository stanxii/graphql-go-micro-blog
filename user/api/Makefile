GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=../srv --micro_out=. --go_out=. proto/user/user.proto

.PHONY: build
build: proto
	go build -o user-api *.go

.PHONY: docker
docker:
	docker build . -t user-api:latest