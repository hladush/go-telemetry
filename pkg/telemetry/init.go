package telemetry

import (
	"fmt"
	"os"

	"github.com/hladush/go-telemetry/internal/log"
	"github.com/hladush/go-telemetry/internal/metrics"
	"github.com/hladush/go-telemetry/internal/utils"
)

const (
	logFmt           = "[%s.%s]"
	metricsFmt       = "%s_%s"
	metricsPrefixFmt = "%s_%s_%s"

	metricsPrefixEnv = "METRICS_PREFIX"
	metricsDefault   = "go_app_telemetry"

	metricsEmitterEnv     = "METRICS_EMITTER"
	defaultMetricsEmitter = "console"

	prometheusPortEnv     = "PROMETHEUS_ADDRESS"
	defaultPrometheusPort = ":8081"
)

var (
	metricsPrefix  = ""
	metricsEmitter MetricsEmitter
	logger         Logger
)

func init() {
	metricsPrefix = utils.GetEnvStringOrDefault(metricsPrefixEnv, metricsDefault)
	initCollector()
	initLogger()
}

func initLogger() {
	logger = &log.ConsoleLogger{}
}

func initCollector() {
	emitter := utils.GetEnvStringOrDefault(metricsEmitterEnv, defaultMetricsEmitter)

	switch emitter {
	case "console":
		metricsEmitter = &metrics.ConsoleMetrics{}
	case "noop":
		metricsEmitter = &metrics.NoopMetrics{}
	case "prometheus":
		address := utils.GetEnvStringOrDefault(prometheusPortEnv, defaultPrometheusPort)
		metricsEmitter = metrics.NewPrometheusMetrics(address)
	default:
		fmt.Printf("Unknown metrics emitter: %s\n", emitter)
		os.Exit(1)
	}
}
