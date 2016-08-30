package main

import (
	"fmt"
	"time"

	metrics "../../metrics"
	metricslib "../../metrics/metrics-lib"
	"github.com/pelletier/go-toml"
)

func ReadConfig() {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		// retrieve data directly
		user := config.Get("postgres.user").(string)
		password := config.Get("postgres.password").(string)

		// or using an intermediate object
		configTree := config.Get("postgres").(*toml.TomlTree)
		user = configTree.Get("user").(string)
		password = configTree.Get("password").(string)
		fmt.Println("User is ", user, ". Password is ", password)

		// show where elements are in the file
		fmt.Println("User position: %v", configTree.GetPosition("user"))
		fmt.Println("Password position: %v", configTree.GetPosition("password"))

		// use a query to gather elements without walking the tree
		results, _ := config.Query("$..[user,password]")
		for ii, item := range results.Values() {
			fmt.Println("Query result %d: %v", ii, item)
		}
	}
}

func Module1(ms *metrics.DMetrics) {
	m1 := make(map[string]interface{})
	m1["module1_counter"] = ms.RegMetric("module1_counter", metricslib.NewCounter())
	m1["module1_action"] = ms.RegMetric("module1_action", metricslib.NewFlow(50))
	for {
		m1["module1_counter"].(metricslib.Counter).Inc(1)
		m1["module1_action"].(metricslib.Flow).RPush(` has been added to venuscron. `)
		time.Sleep(3 * time.Second)
	}
}

func Module2(ms *metrics.DMetrics) {
	m2 := make(map[string]interface{})
	m2["module2_counter"] = ms.RegMetric("module2_counter", metricslib.NewCounter())
	m2["module2_action"] = ms.RegMetric("module2_action", metricslib.NewFlow(50))
	for {
		m2["module2_counter"].(metricslib.Counter).Inc(1)
		m2["module2_action"].(metricslib.Flow).RPush(` has been added to venuscron. `)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	redisPools = newRdsPool(serverConf.ESWrite, serverConf.ESWriteAuth)

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	ms := metrics.NewDMetrics(&metrics.Options{
		Switcher: 1,
		Path:     "",
		Addr:     "127.0.0.1:9090",
	})
	go Module1(ms)
	go Module2(ms)
	time.Sleep(3 * time.Second)

	for key, value := range ms.Metrics {
		conn := redisPools.Get()
		RPush(conn, key, value)
	}
}
