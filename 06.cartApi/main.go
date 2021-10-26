package main

import (
	"github.com/gzltommy/cart/proto/cart"
	"github.com/gzltommy/cartApi/client"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	"net"
	"net/http"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/micro/go-micro/v2"
	"github.com/gzltommy/cartApi/handler"
	cartApi "github.com/gzltommy/cartApi/proto/cartApi"
	"github.com/micro/go-micro/v2/registry"
	"github.com/gzltommy/common"
)

func main() {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			fmt.Sprintf("%s:%d", common.Consul_Host, 8500),
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.cartApi", fmt.Sprintf("%s:%d", common.Tracer_Host, 6831))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("", "8000"), hystrixStreamHandler)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		//micro.Address("localhost:8087"),

		//添加 consul 作为注册中心
		micro.Registry(consulRegistry),

		// 给提供的服务（不是客户端），绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),

		// 添加熔断
		micro.WrapClient(client.NewClientHystrixWrapper()),

		// 添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),


		//// 添加限流
		//micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Initialise service
	service.Init()

	cartService := cart.NewCartService("go.micro.service.cart", service.Client())

	// Register Handler
	err = cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService: cartService})
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
