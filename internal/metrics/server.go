package metrics

import (
	"github.com/containerssh/libcontainerssh/config"
	http2 "github.com/containerssh/libcontainerssh/http"
	"github.com/containerssh/libcontainerssh/log"
	messageCodes "github.com/containerssh/libcontainerssh/message"
)

// NewServer creates a new metrics server based on the configuration. It MAY return nil if the server is disabled.
func NewServer(cfg config.MetricsConfig, collector Collector, logger log.Logger) (http2.Server, error) {
	if !cfg.Enable {
		return nil, nil
	}
	return http2.NewServer(
		"Metrics server",
		cfg.HTTPServerConfiguration,
		NewHandler(
			cfg.Path,
			collector,
		),
		logger,
		func(url string) {
			logger.Info(
				messageCodes.NewMessage(
					messageCodes.MHealthServiceAvailable,
					"Metrics server is now available at %s%s",
					url, cfg.Path,
				))
		},
	)
}