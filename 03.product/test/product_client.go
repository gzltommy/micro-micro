package main

import (
	"context"
	"github.com/gzltommy/product/common"
	"github.com/gzltommy/product/proto/product"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/registry"
	"fmt"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			fmt.Sprintf("%s:%d", common.Consul_Host, 8500),
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.product.client", fmt.Sprintf("%s:%d", common.Tracer_Host, 6831))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 创建服务
	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		//micro.Address("127.0.0.1:8086"),

		//添加 consul 作为注册中心
		micro.Registry(consulRegistry),

		// 给客户端，绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	// 初始化服务
	service.Init()

	/*--------------------------------------------------------------------------------------------------------------*/
	/*                                                核心代码区域                                                  */
	/*--------------------------------------------------------------------------------------------------------------*/
	// 将该服务作为客户端与 "go.micro.service.product" 进行连接，获取到 "go.micro.service.product" 服务实例
	productService := product.NewProductService("go.micro.service.product", service.Client())

	// 调用 "go.micro.service.product" 服务提供的方法
	response, err := productService.AddProduct(context.TODO(), &product.ProductInfo{
		//Id:                 0,
		ProductName:        "小米手机",
		ProductSku:         "xiaomi_phone_001",
		ProductPrice:       1998,
		ProductDescription: "好用不贵",
		ProductCategoryId:  1,
		ProductImage: []*product.ProductImage{
			{
				ImageName: "xiaomi_phone_image01",
				ImageCode: "xiaomi_phone_image01_code",
				ImageUrl:  "xiaomi_phone_image01_url",
			},
			{
				ImageName: "xiaomi_phone_image02",
				ImageCode: "xiaomi_phone_image02_code",
				ImageUrl:  "xiaomi_phone_image02_url",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "xiaomi_phone_size1",
				SizeCode: "xiaomi_phone_size1_code",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "xiaomi_phone_seo",
			SeoKeywords:    "xiaomi_phone_seo",
			SeoDescription: "xiaomi_phone_seo",
			SeoCode:        "xiaomi_phone_seo",
		},
	})

	/*--------------------------------------------------------------------------------------------------------------*/
	/*                                                核心代码区域                                                  */
	/*--------------------------------------------------------------------------------------------------------------*/

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

	// 开启服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
