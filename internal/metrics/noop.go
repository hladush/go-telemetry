package metrics

type NoopMetrics struct{}

func (n *NoopMetrics) IncCounter(metric string)             {}
func (n *NoopMetrics) Observe(metric string, value float64) {}
