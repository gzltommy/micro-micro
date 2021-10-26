module github.com/gzltommy/cartApi

go 1.13

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/golang/protobuf v1.5.2
	github.com/gzltommy/cart v0.0.0-20210729080354-2b2168f87e3a
	github.com/gzltommy/common v0.0.0-20210726061703-2309305f3bee
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.9.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
