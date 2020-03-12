GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto-article:
	protoc --proto_path=../srv --micro_out=. --go_out=. proto/article/article.proto

proto: proto-article;

.PHONY: build
build: proto 
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o article-api server.go

.PHONY: docker
docker:
	docker build . -t article-api:latest
