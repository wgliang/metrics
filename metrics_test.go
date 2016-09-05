package metrics

import (
	"fmt"
	"testing"
	"time"
)

func Module1(ms *Metrics) {
	m1 := make(map[string]interface{})
	m1["module1_counter"] = ms.RegMetric("module1_counter", NewCounter())
	m1["module1_action"] = ms.RegMetric("module1_action", NewFlow(50))
	for i := 0; i < 3; i++ {
		m1["module1_counter"].(Counter).Inc(1)
		m1["module1_action"].(Flow).RPush(` has been added to venuscron. `)
		// time.Sleep(3 * time.Second)
	}
}

func Module2(ms *Metrics) {
	m2 := make(map[string]interface{})
	m2["module2_counter"] = ms.RegMetric("module2_counter", NewCounter())
	m2["module2_action"] = ms.RegMetric("module2_action", NewFlow(50))
	for i := 0; i < 3; i++ {
		m2["module2_counter"].(Counter).Inc(1)
		m2["module2_action"].(Flow).RPush(` has been added to venuscron. `)
		// time.Sleep(3 * time.Second)
	}
}

func Test_Metrics(t *testing.T) {
	ms := NewMetrics(&Options{
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
