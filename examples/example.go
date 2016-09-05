package main

import (
	"fmt"
	"time"

	metrics "github.com/wgliang/metrics"
)

func Module1(ms *metrics.Metrics) {
	m1 := make(map[string]interface{})
	m1["module1_counter"] = ms.RegMetric("module1_counter", metrics.NewCounter())
	m1["module1_action"] = ms.RegMetric("module1_action", metrics.NewFlow(50))
	for i := 0; i < 3; i++ {
		m1["module1_counter"].(metrics.Counter).Inc(1)
		m1["module1_action"].(metrics.Flow).RPush(` has been added to venuscron. `)
		// time.Sleep(3 * time.Second)
	}
}

func Module2(ms *metrics.Metrics) {
	m2 := make(map[string]interface{})
	m2["module2_counter"] = ms.RegMetric("module2_counter", metrics.NewCounter())
	m2["module2_action"] = ms.RegMetric("module2_action", metrics.NewFlow(50))
	for i := 0; i < 3; i++ {
		m2["module2_counter"].(metrics.Counter).Inc(1)
		m2["module2_action"].(metrics.Flow).RPush(` has been added to venuscron. `)
		// time.Sleep(3 * time.Second)
	}
}

func main() {
	ms := metrics.NewMetrics(&metrics.Options{
		Switcher: 1,
		Path:     "",
		Addr:     "127.0.0.1:9090",
	})
	go Module1(ms)
	go Module2(ms)
	time.Sleep(3 * time.Second)

	for key, _ := range ms.Metrics {
		value, err := ms.Metrics2string(key)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(key + " , " + value)
	}
}
