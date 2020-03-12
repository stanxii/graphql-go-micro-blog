
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --micro_out=. --go_out=. proto/user/user.proto

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o user-srv cmd/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-srv:latest
