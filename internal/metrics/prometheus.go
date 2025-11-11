package metrics

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusMetrics struct {
	counterByName map[string]prometheus.Counter
	counterRWMux  sync.RWMutex

	summaryByName map[string]prometheus.Summary
	summaryRWMux  sync.RWMutex
	
	gaugeByName   map[string]prometheus.Gauge
	gaugeRWMux    sync.RWMutex
}

func NewPrometheusMetrics(address string) *PrometheusMetrics {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	go func() {
		if err := http.ListenAndServe(address, mux); err != nil {
			panic(fmt.Sprintf("Failed to start Prometheus metrics server: %v", err))
		}
	}()

	return &PrometheusMetrics{
		counterByName: make(map[string]prometheus.Counter),
		counterRWMux:  sync.RWMutex{},

		summaryByName: make(map[string]prometheus.Summary),
		summaryRWMux:  sync.RWMutex{},

		gaugeByName:   make(map[string]prometheus.Gauge),
		gaugeRWMux:    sync.RWMutex{},
	}
}

func (p *PrometheusMetrics) Separator() string {
	return "_"
}
func (p *PrometheusMetrics) IncCounter(metric string) {
	p.createCounterIfDoesntExist(metric)
	p.counterByName[metric].Inc()

}

func (p *PrometheusMetrics) Observe(metric string, value float64) {
	p.createSummaryIfDoesntExist(metric)
	p.summaryByName[metric].Observe(value)
}

func (p *PrometheusMetrics) createSummaryIfDoesntExist(metric string) {
	p.summaryRWMux.RLock()
	_, exists := p.summaryByName[metric]
	p.summaryRWMux.RUnlock()
	if exists {
		return
	}

	p.summaryRWMux.Lock()
	defer p.summaryRWMux.Unlock()
	if _, exists := p.summaryByName[metric]; exists {
		return
	}
	summary := prometheus.NewSummary(prometheus.SummaryOpts{
		Name: metric,
		Objectives: map[float64]float64{
			0.5:  0.01,
			0.9:  0.01,
			0.95: 0.01,
			0.99: 0.005,
		},
	})
	p.summaryByName[metric] = summary
	prometheus.MustRegister(summary)
}

func (p *PrometheusMetrics) SetGauge(metric string, value float64) {
	p.createGaugeIfDoesntExist(metric)
	p.gaugeByName[metric].Set(value)
}

func (p *PrometheusMetrics) createGaugeIfDoesntExist(metric string) {
	p.gaugeRWMux.RLock()
	_, exists := p.gaugeByName[metric]
	p.gaugeRWMux.RUnlock()
	if exists {
		return
	}

	p.gaugeRWMux.Lock()
	defer p.gaugeRWMux.Unlock()
	if _, exists := p.gaugeByName[metric]; exists {
		return
	}
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: metric,
	})
	p.gaugeByName[metric] = gauge
	prometheus.MustRegister(gauge)
}

func (p *PrometheusMetrics) createCounterIfDoesntExist(metric string) {
	p.counterRWMux.RLock()
	_, exists := p.counterByName[metric]
	p.counterRWMux.RUnlock()
	if exists {
		return
	}

	p.counterRWMux.Lock()
	defer p.counterRWMux.Unlock()
	if _, exists := p.counterByName[metric]; exists {
		return
	}
	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: metric,
		},
	)
	p.counterByName[metric] = counter
	prometheus.MustRegister(counter)
}
