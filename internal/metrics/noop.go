package metrics

type NoopMetrics struct{}

func (n *NoopMetrics) Separator() string {
	return ""
}
func (n *NoopMetrics) IncCounter(metric string)                  {}
func (n *NoopMetrics) IncCounterBy(metric string, value float64) {}
func (n *NoopMetrics) Observe(metric string, value float64)      {}
func (n *NoopMetrics) SetGauge(metric string, value float64)     {}
