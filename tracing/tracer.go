package tracing

import (
	"time"
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"

)

var (
	defaultSampleRatio float64 = 1
)

// Init returns a newly configured tracer
func Init(serviceName, host string) (opentracing.Tracer, error) {
	ratio := defaultSampleRatio
	log.Printf("jaeger: tracing sample ratio %f", ratio)
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "probabilistic",
			Param: ratio,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  host,
		},
	}
	logger := jaegerlog.StdLogger
	tracer, _, err := cfg.New(serviceName, jaegercfg.Logger(logger))
	if err != nil {
		return nil, err
	}
	return tracer, nil
}
