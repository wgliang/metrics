package metrics

import (
	"fmt"
	"reflect"
	"sync"
)

// DuplicateMetric is the error returned by Collectry.Collector when a metric
// already exists.  If you mean to Collector that metric you must first
// UnCollector the existing metric.
type DuplicateMetric string

func (err DuplicateMetric) Error() string {
	return fmt.Sprintf("duplicate metric: %s", string(err))
}

// A Collectry holds references to a set of metrics by name and can iterate
// over them, calling callback functions provided by the user.
//
// This is an interface so as to encourage other structs to implement
// the Collectry API as appropriate.
type Collectry interface {

	// Call the given function for each Collectored metric.
	Each(func(string, interface{}))

	// Get the metric by the given name or nil if none is Collectored.
	Get(string) interface{}

	// Gets an existing metric or Collectors the given one.
	// The interface can be the metric to Collector if not found in Collectry,
	// or a function returning the metric for lazy instantiation.
	GetOrCollector(string, interface{}) interface{}

	// Collector the given metric under the given name.
	Collector(string, interface{}) error

	// Run all Collectored healthchecks.
	RunHealthchecks()

	// UnCollector the metric with the given name.
	Uncollector(string)

	// UnCollector all metrics.  (Mostly for testing.)
	UncollectorAll()

	// Get metrics
	Values() interface{}

	// Get metrics number
	Size() int64
}

// The standard implementation of a Collectry is a mutex-protected map
// of names to metrics.
type StandardCollectry struct {
	metrics map[string]interface{}
	mutex   sync.RWMutex
}

// Create a new Collectry.
func NewCollectry() Collectry {
	return &StandardCollectry{metrics: make(map[string]interface{})}
}

// Call the given function for each Collectored metric.
func (r *StandardCollectry) Each(f func(string, interface{})) {
	r.mutex.RLock()
	cs := r.collectored()
	r.mutex.RUnlock()
	for name, i := range cs {
		f(name, i)
	}
}

// Get the metric by the given name or nil if none is Collectored.
func (r *StandardCollectry) Get(name string) interface{} {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.metrics[name]
}

// Gets an existing metric or creates and Collectors a new one. Threadsafe
// alternative to calling Get and Collector on failure.
// The interface can be the metric to Collector if not found in Collectry,
// or a function returning the metric for lazy instantiation.
func (r *StandardCollectry) GetOrCollector(name string, i interface{}) interface{} {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if metric, ok := r.metrics[name]; ok {
		return metric
	}
	if v := reflect.ValueOf(i); v.Kind() == reflect.Func {
		i = v.Call(nil)[0].Interface()
	}
	r.collector(name, i)
	return i
}

// Collector the given metric under the given name.  Returns a DuplicateMetric
// if a metric by the given name is already Collectored.
func (r *StandardCollectry) Collector(name string, i interface{}) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.collector(name, i)
}

// Run all Collectored healthchecks.
func (r *StandardCollectry) RunHealthchecks() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, i := range r.metrics {
		if h, ok := i.(Healthcheck); ok {
			h.Check()
		}
	}
}

// UnCollector the metric with the given name.
func (r *StandardCollectry) Uncollector(name string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.metrics, name)
}

// UnCollector all metrics.  (Mostly for testing.)
func (r *StandardCollectry) UncollectorAll() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for name, _ := range r.metrics {
		delete(r.metrics, name)
	}
}

// size of my metrics
func (r *StandardCollectry) Size() int64 {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return int64(len(r.metrics))
}

func (r *StandardCollectry) collector(name string, i interface{}) error {
	if _, ok := r.metrics[name]; ok {
		return DuplicateMetric(name)
	}
	switch i.(type) {
	case Counter, Gauge, GaugeFloat64, Healthcheck, Histogram, Meter, Timer:
		r.metrics[name] = i
	}
	return nil
}

func (r *StandardCollectry) collectored() map[string]interface{} {
	metrics := make(map[string]interface{}, len(r.metrics))
	for name, i := range r.metrics {
		metrics[name] = i
	}
	return metrics
}

type PrefixedCollectry struct {
	underlying Collectry
	prefix     string
}

func NewPrefixedCollectry(prefix string) Collectry {
	return &PrefixedCollectry{
		underlying: NewCollectry(),
		prefix:     prefix,
	}
}

func NewPrefixedChildCollectry(parent Collectry, prefix string) Collectry {
	return &PrefixedCollectry{
		underlying: parent,
		prefix:     prefix,
	}
}

// Call the given function for each Collectored metric.
func (r *PrefixedCollectry) Each(fn func(string, interface{})) {
	r.underlying.Each(fn)
}

// Get the metric by the given name or nil if none is Collectored.
func (r *PrefixedCollectry) Get(name string) interface{} {
	return r.underlying.Get(name)
}

// Gets an existing metric or Collectors the given one.
// The interface can be the metric to Collector if not found in Collectry,
// or a function returning the metric for lazy instantiation.
func (r *PrefixedCollectry) GetOrCollector(name string, metric interface{}) interface{} {
	realName := r.prefix + name
	return r.underlying.GetOrCollector(realName, metric)
}

// Collector the given metric under the given name. The name will be prefixed.
func (r *PrefixedCollectry) Collector(name string, metric interface{}) error {
	realName := r.prefix + name
	return r.underlying.Collector(realName, metric)
}

// Run all Collectored healthchecks.
func (r *PrefixedCollectry) RunHealthchecks() {
	r.underlying.RunHealthchecks()
}

// UnCollector the metric with the given name. The name will be prefixed.
func (r *PrefixedCollectry) Uncollector(name string) {
	realName := r.prefix + name
	r.underlying.Uncollector(realName)
}

// UnCollector all metrics.  (Mostly for testing.)
func (r *PrefixedCollectry) UncollectorAll() {
	r.underlying.UncollectorAll()
}

func (r *PrefixedCollectry) Values() interface{} {
	return r.underlying.Values()
}

// size of my metrics
func (r *PrefixedCollectry) Size() int64 {
	return r.underlying.Size()
}

var DefaultCollectry Collectry = NewCollectry()

// Call the given function for each Collectored metric.
func Each(f func(string, interface{})) {
	DefaultCollectry.Each(f)
}

// Get the metric by the given name or nil if none is Collectored.
func Get(name string) interface{} {
	return DefaultCollectry.Get(name)
}

// Gets an existing metric or creates and Collectors a new one. Threadsafe
// alternative to calling Get and Collector on failure.
func GetOrCollector(name string, i interface{}) interface{} {
	return DefaultCollectry.GetOrCollector(name, i)
}

// Collector the given metric under the given name.  Returns a DuplicateMetric
// if a metric by the given name is already Collectored.
func Collector(name string, i interface{}) error {
	return DefaultCollectry.Collector(name, i)
}

// Collector the given metric under the given name.  Panics if a metric by the
// given name is already Collectored.
func MustCollector(name string, i interface{}) {
	if err := Collector(name, i); err != nil {
		panic(err)
	}
}

// Run all Collectored healthchecks.
func RunHealthchecks() {
	DefaultCollectry.RunHealthchecks()
}

// UnCollector the metric with the given name.
func UnCollector(name string) {
	DefaultCollectry.Uncollector(name)
}

// Get metrics number
func Size() int64 {
	return DefaultCollectry.Size()
}

func (r *StandardCollectry) Values() interface{} {
	data := make(map[string]map[string]interface{})
	r.Each(func(name string, i interface{}) {
		values := make(map[string]interface{})
		switch metric := i.(type) {
		case Counter:
			values["count"] = metric.Count()
		case Gauge:
			values["value"] = metric.Value()
		case GaugeFloat64:
			values["value"] = metric.Value()
		case Healthcheck:
			values["error"] = nil
			metric.Check()
			if err := metric.Error(); nil != err {
				values["error"] = metric.Error().Error()
			}
		case Meter:
			// m := metric.Snapshot()
			values["sum"] = metric.Count()
			// values["rate"] = metric.RateStep()
			values["rate"] = 0.0
			values["rate.1min"] = metric.Rate1()
			values["rate.5min"] = metric.Rate5()
			values["rate.15min"] = metric.Rate15()
		case Histogram:
			// h := metric.Snapshot()
			ps := metric.Percentiles([]float64{0.5, 0.75, 0.99})
			values["min"] = metric.Min()
			values["max"] = metric.Max()
			values["mean"] = metric.Mean()
			values["75th"] = ps[0]
			values["95th"] = ps[1]
			values["99th"] = ps[2]
		case Timer:
			// t := metric.Snapshot()
			ps := metric.Percentiles([]float64{0.5, 0.75, 0.99})
			values["min"] = metric.Min()
			values["max"] = metric.Max()
			values["mean"] = metric.Mean()
			values["75th"] = ps[0]
			values["95th"] = ps[1]
			values["99th"] = ps[2]
			values["sum"] = metric.Count()
			// values["rate"] = metric.RateStep()
			values["rate"] = 0.0
			values["rate.1min"] = metric.Rate1()
			values["rate.5min"] = metric.Rate5()
			values["rate.15min"] = metric.Rate15()
		}
		data[name] = values
	})
	return data
}
