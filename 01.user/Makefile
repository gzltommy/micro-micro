
GOPATH:=$(shell go env GOPATH)
# 在命令行中使用 M 标志，指定 .proto 文件生成的 .pb.go 文件的包名（含路径）
# 在容器中执行 protoc 命令要使用规范的形式指定包路径：
# ① 在 .proto 文件中使用 go_package 选项指定（推荐只指定路径，不指定包名，包名由路径推演生成）；
# ② 在调用 protoc 编译器时使用 M 标记在选项中指定（优先级高于①）
# MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto
# MODIFY=Mproto/user/user.proto=proto/user

# 获取当前工作目录
PWD=$(shell pwd)

APPNAME=user

.PHONY: proto
proto:
	@#protoc3 --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/user/user.proto
	@#docker run --rm -v $(pwd):$(pwd) -w $(pwd) gzltommy/protoc   -I ./ --go_out=./   --micro_out=./   ./*.proto	
	@docker run --rm -v $(PWD):$(PWD) -w $(PWD) gzltommy/protoc --proto_path=. --micro_out=paths=source_relative:. --go_out=paths=source_relative:. proto/user/user.proto

.PHONY: build
build: proto
	go build -o $(APPNAME) *.go
# 因为权限问题无法解决（挂载的windows上的共享文件夹），所有不能用下面的方式	
#	docker run --rm -v $(PWD):$(PWD) -v $(GOPATH):$(GOPATH) -w $(PWD) -e GOPATH=$(PWD) -e GOOS=linux -e GOARCH=amd64 gzltommy/golang go build -o $(APPNAME) *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t $(APPNAME):latest
