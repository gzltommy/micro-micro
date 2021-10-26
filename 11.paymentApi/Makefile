
GOPATH:=$(shell go env GOPATH)
#MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto


# 获取当前工作目录
PWD=$(shell pwd)

APPNAME=paymentApi

.PHONY: proto
proto:
	@docker run --rm -v $(PWD):$(PWD) -w $(PWD) gzltommy/protoc --proto_path=. --micro_out=paths=source_relative:. --go_out=paths=source_relative:. proto/paymentApi/*.proto
	@docker run --rm -v $(PWD):$(PWD) -w $(PWD) gzltommy/protoc --proto_path=. --micro_out=paths=source_relative:. --go_out=paths=source_relative:. proto/imports/*.proto
#	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/cartApi/cartApi.proto
    

.PHONY: build
build: proto
	go build -o $(APPNAME) *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t $(APPNAME):latest
