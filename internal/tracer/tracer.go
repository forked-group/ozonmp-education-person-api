package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"io"

	"github.com/aaa2ppp/ozonmp-education-kw-person-api/internal/config"
	"github.com/rs/zerolog/log"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer - returns new tracer.
func NewTracer(cfg *config.Config) (io.Closer, error) {
	cfgTracer := &jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.Service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Jaeger.Host + cfg.Jaeger.Port,
		},
	}
	tracer, closer, err := cfgTracer.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		log.Err(err).Msgf("failed init jaeger: %v", err)

		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	log.Info().Msgf("Traces started")

	return closer, nil
}
