package main

import (
	"github.com/gzltommy/payment/proto/payment"
	"github.com/gzltommy/paymentApi/client"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	"net"
	"net/http"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/micro/go-micro/v2"
	"github.com/gzltommy/paymentApi/handler"
	paymentApi "github.com/gzltommy/paymentApi/proto/paymentApi"
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
	t, io, err := common.NewTracer("go.micro.service.paymentApi", fmt.Sprintf("%s:%d", common.Tracer_Host, 6831))
	if err != nil {
		common.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("", "8000"), hystrixStreamHandler)
		if err != nil {
			common.Fatal(err)
		}
	}()

	common.PrometheusBoot("9292")

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.paymentApi"),
		micro.Version("latest"),
		//micro.Address("localhost:8087"),

		//添加 consul 作为注册中心
		micro.Registry(consulRegistry),

		// 链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),

		// 添加熔断
		micro.WrapClient(client.NewClientHystrixWrapper()),

		// 添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),


		//// 添加限流
		//micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Initialise service
	service.Init()

	paymentService := payment.NewPaymentService("go.micro.service.payment", service.Client())

	// Register Handler
	err = paymentApi.RegisterPaymentApiHandler(service.Server(), &handler.PaymentApi{PaymentService: paymentService})
	if err != nil {
		common.Fatal(err)
	}

	// Run service
	if err = service.Run(); err != nil {
		common.Fatal(err)
	}
}
