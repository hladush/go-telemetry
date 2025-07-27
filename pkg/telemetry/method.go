package telemetry

import (
	"fmt"
	"time"

	"github.com/hladush/go-telemetry/internal/utils"
)

type Method struct {
	logName       string
	metricsPrefix string
	errorPrefix   string
	successPrefix string
	requestPrefix string
	latencyPrefix string
}

type MetricsEmiter interface {
	IncCounter(metric string)
	Observe(metric string, value float64)
}

type Logger interface {
	LogInfo(message string)
	LogDebug(message string)
	LogError(message string)
}

func NewMethod(methodName, serviceName string) *Method {
	snakeCaseServiceName := utils.ToSnakeCase(serviceName)
	snakeCaseMethodName := utils.ToSnakeCase(methodName)
	return &Method{
		logName:       fmt.Sprintf(logFmt, serviceName, methodName),
		metricsPrefix: getMetricsName(methodName, serviceName),
		errorPrefix:   getMetricsName(fmt.Sprintf("%s.error", snakeCaseMethodName), snakeCaseServiceName),
		successPrefix: getMetricsName(fmt.Sprintf("%s.success", snakeCaseMethodName), snakeCaseServiceName),
		requestPrefix: getMetricsName(fmt.Sprintf("%s.request", snakeCaseMethodName), snakeCaseServiceName),
		latencyPrefix: getMetricsName(fmt.Sprintf("%s.latency", snakeCaseMethodName), snakeCaseServiceName),
	}
}

func (m *Method) RecordLatency(startTime time.Time, dimension ...string) {
	metric := utils.JoinWithPrefix(m.latencyPrefix, dimension...)
	metricsEmitter.Observe(metric, float64(time.Since(startTime).Milliseconds()))
}

func (m *Method) LogAndCountErrorOrSuccess(err error, dimension ...string) {
	if err != nil {
		m.LogAndCountError(err, dimension...)
		return

	}
	m.LogAndCountSuccess(dimension...)

}

func (m *Method) LogAndCountSuccess(dimension ...string) {
	logger.LogDebug(fmt.Sprintf("%s finished successfuly.| dimension: %v", m.logName, dimension))
	m.CountSuccess(dimension...)
}

func (m *Method) LogAndCountError(err error, dimension ...string) {
	logger.LogError(fmt.Sprintf("%s failed. Reason: %v | dimension: %v", m.logName, err, dimension))
	m.CountError(dimension...)
}

func (m *Method) CountRequest(dimension ...string) {
	metric := utils.JoinWithPrefix(m.requestPrefix, dimension...)
	metricsEmitter.IncCounter(metric)
}

func (m *Method) CountError(dimension ...string) {
	metric := utils.JoinWithPrefix(m.errorPrefix, dimension...)
	metricsEmitter.IncCounter(metric)
}

func (m *Method) CountSuccess(dimension ...string) {
	metric := utils.JoinWithPrefix(m.successPrefix, dimension...)
	metricsEmitter.IncCounter(metric)
}

func (m *Method) IncCounter(dimension ...string) {
	metric := utils.JoinWithPrefix(m.metricsPrefix, dimension...)
	metricsEmitter.IncCounter(metric)
}

func getMetricsName(methodName, serviceName string) string {
	if metricsPrefix == "" {
		return fmt.Sprintf("%s.%s", utils.ToSnakeCase(serviceName), utils.ToSnakeCase(methodName))
	}

	return fmt.Sprintf(metricsPrefixFmt, metricsPrefix, utils.ToSnakeCase(serviceName), utils.ToSnakeCase(methodName))
}
