package common

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

// 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 设置配置中心的地址
		consul.WithAddress(host+":"+fmt.Sprintf("%v", port)),
		// 设置前缀，不设置默认前缀 /micro/config
		consul.WithPrefix(prefix),
		//是否移除前缀，这里是设置为 true，表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)

	// 配置初始化
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	//加载配置
	err = cfg.Load(consulSource)
	return cfg, err
}
