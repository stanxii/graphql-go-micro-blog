GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto-article:
	protoc --proto_path=../srv --micro_out=. --go_out=. proto/article/article.proto

proto: proto-article;

.PHONY: build
build: proto 
	go build -o article-api *.go

.PHONY: docker
docker:
	docker build . -t article-api:latest
