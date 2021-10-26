package main

import (
	"fmt"
	"github.com/gzltommy/product/common"
	"github.com/gzltommy/product/domain/repository"
	productService "github.com/gzltommy/product/domain/service"
	"github.com/gzltommy/product/handler"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing "github.com/opentracing/opentracing-go"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	product "github.com/gzltommy/product/proto/product"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig(common.Consul_Host, 8500, "/micro/config")
	if err != nil {
		panic(err)
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			fmt.Sprintf("%s:%d", common.Consul_Host, 8500),
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.product", fmt.Sprintf("%s:%d", common.Tracer_Host, 6831))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 创建数据库连接
	// 获取 mysql 配置，路径中不带前缀
	mysqlConfig, err := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	if err != nil {
		panic(err)
	}

	//dsn := "root:123456@tcp(192.168.100.64:3307)/micro2?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Pwd, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 只会执行一次，数据表的自动迁移（创建）
	rp := repository.NewProductRepository(db)
	rp.InitTable()

	// 创建服务实例
	dataService := productService.NewProductDataService(rp)

	// 创建服务
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		//micro.Address("127.0.0.1:8085"),

		//添加 consul 作为注册中心
		micro.Registry(consulRegistry),

		// 给提供的服务（不是客户端），绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// 初始化服务
	service.Init()

	// 注册 Handler
	product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService: dataService})

	// 开启服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
