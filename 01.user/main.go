package main

import (
	"github.com/gzltommy/user/domain/repository"
	userService "github.com/gzltommy/user/domain/service"
	"github.com/gzltommy/user/handler"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	user "github.com/gzltommy/user/proto/user"
)

func main() {
	// 创建服务
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	// 初始化服务
	service.Init()

	// 创建数据库连接
	dsn := "root:123456@tcp(192.168.100.64:3307)/micro2?charset=utf8mb4&parseTime=True&loc=Local"
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
	rp := repository.NewUserRepository(db)
	rp.InitTable()
	
	// 创建服务实例
	userDataService := userService.NewUserDataService(rp)

	// 注册 Handler
	user.RegisterUserHandler(service.Server(), &handler.User{UserDataService: userDataService})

	// 开启服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
