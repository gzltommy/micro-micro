package common

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

// 创建链路追踪实例
func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
			//SamplingServerURL:        "",
			//SamplingRefreshInterval:  0,
			//MaxOperations:            0,
			//OperationNameLateBinding: false,
			//Options:                  nil,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: 1 * time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  addr,
			//QueueSize:                  0,
			//DisableAttemptReconnecting: false,
			//AttemptReconnectInterval:   0,
			//CollectorEndpoint:          "",
			//User:                       "",
			//Password:                   "",
			//HTTPHeaders:                nil,
		},
		//Disabled:    false,
		//RPCMetrics:  false,
		//Gen128Bit:   false,
		//Tags:        nil,
		//Headers:             nil,
		//BaggageRestrictions: nil,
		//Throttler:           nil,
	}
	return cfg.NewTracer()
}
