# Metrics

[![Go Report Card](https://travis-ci.org/wgliang/metrics.svg?branch=master)](https://travis-ci.org/wgliang/metrics)
[![GoDoc](https://godoc.org/github.com/wgliang/metrics?status.svg)](https://godoc.org/github.com/wgliang/metrics)
[![Go Report Card](https://goreportcard.com/badge/github.com/wgliang/metrics)](https://goreportcard.com/report/github.com/wgliang/metrics)

forked from https://github.com/facebookgo/metrics and added some useful types


## Introduction

When I use facebook/metrics I find the type of it support just for int64 or int64[], but I should take all the events and error-logs in my developed modular. So I add the flow(流) type and mstring,then you can take some type-string data into it.

## How to use it ?

```
package main

import (
	"fmt"
	"time"

	metrics "github.com/wgliang/metrics"
	metricslib "github.com/wgliang/metrics/metrics-lib"
)

func Module1(ms *metrics.Metrics) {
	m1 := make(map[string]interface{})
	m1["module1_counter"] = ms.RegMetric("module1_counter", metricslib.NewCounter())
	m1["module1_action"] = ms.RegMetric("module1_action", metricslib.NewFlow(50))
	for i := 0; i < 3; i++ {
		m1["module1_counter"].(metricslib.Counter).Inc(1)
		m1["module1_action"].(metricslib.Flow).RPush(` has been added to venuscron. `)
		// time.Sleep(3 * time.Second)
	}
}

func Module2(ms *metrics.Metrics) {
	m2 := make(map[string]interface{})
	m2["module2_counter"] = ms.RegMetric("module2_counter", metricslib.NewCounter())
	m2["module2_action"] = ms.RegMetric("module2_action", metricslib.NewFlow(50))
	for i := 0; i < 3; i++ {
		m2["module2_counter"].(metricslib.Counter).Inc(1)
		m2["module2_action"].(metricslib.Flow).RPush(` has been added to venuscron. `)
		// time.Sleep(3 * time.Second)
	}
}

func main() {
	ms := metrics.NewMetrics(&metrics.Options{
		Switcher: 1,
		Path:     "{static-files}",
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

```

## Or..

see data from 127.0.0.1:9090/api/metrics/metrics/{action} , "action" is the name that you have register in modules.
