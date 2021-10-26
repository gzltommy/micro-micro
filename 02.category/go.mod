module github.com/gzltommy/category

go 1.13

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/config/source/consul/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
