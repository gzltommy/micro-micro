module github.com/gzltommy/cart

go 1.13

require (
	github.com/golang/protobuf v1.4.3
	github.com/gzltommy/common v0.0.0-20210726061703-2309305f3bee
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/config/source/consul/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
