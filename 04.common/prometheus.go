package _4_common

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

func PrometheusBoot(port string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+port, nil)
		if err != nil {
			log.Fatal("启动失败")
		}
		log.Infof("监控启动，端口为：%s \n", port)
	}()
}
