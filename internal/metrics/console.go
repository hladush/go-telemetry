package metrics

import (
	"fmt"
)

type ConsoleMetrics struct{}

func (c *ConsoleMetrics) IncCounter(metric string) {
	fmt.Printf("Metric: %s, Value: %d\n", metric, 1)
}

func (c *ConsoleMetrics) Observe(metric string, value float64) {
	fmt.Printf("Observe Metric: %s, Value: %f\n", metric, value)
}

func (c *ConsoleMetrics) SetGauge(metric string, value float64) {
	fmt.Printf("Set Gauge Metric: %s, Value: %f\n", metric, value)
}