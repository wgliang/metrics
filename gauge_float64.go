package metrics

import "sync"

// Gauges hold an int64 value that can be set arbitrarily.
type GaugeFloat64 interface {
	// Update the gauge's value.
	Update(value float64)

	// Return the gauge's current value.
	Value() float64
}

// The standard implementation of a Gauge uses the sync/atomic package
// to manage a single int64 value.
type gaugefloat64 struct {
	mutex sync.Mutex
	value float64
}

// Create a new gauge.
func NewGaugeFloat64() GaugeFloat64 {
	return &gaugefloat64{value: 0.0}
}

func (g *gaugefloat64) Update(v float64) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.value = v
}

func (g *gaugefloat64) Value() float64 {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	return g.value
}
