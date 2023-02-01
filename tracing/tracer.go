package tracing

import (
	"time"
	"os"
	"strconv"
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

var (
	defaultSampleRatio float64 = 1
)

// Init returns a newly configured tracer
func Init(serviceName, host string) (opentracing.Tracer, error) {
	ratio := defaultSampleRatio
	log.printf("jaeger at host %s: tracing sample ratio %f", host, ratio)
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "probabilistic",
			Param: ratio,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  host,
		},
	}

	tracer, _, err := cfg.New(serviceName)
	if err != nil {
		return nil, err
	}
	return tracer, nil
}
