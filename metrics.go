package metrics

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"

	"github.com/google/glog"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	metricslib "../../metrics/metrics-lib"
)

// metrics-data options
type Options struct {
	Switcher int
	Path     string
	Addr     string
}

// metrics-data struct
type DMetrics struct {
	mutex   sync.RWMutex
	Metrics map[string]interface{}
	Path    string
	Addr    string
}

// init metrics-data colloctor
func InitDMetrics(me *DMetrics) {
	mux := mux.NewRouter()
	me.RegHandler(mux)

	n := negroni.Classic()
	selfDir, _ := filepath.Abs(me.Path)
	n.Use(negroni.NewStatic(http.Dir(selfDir)))
	n.UseHandler(mux)
	glog.V(5).Infoln(` metrics strating... `)
	//start http server
	go func(serverAddr string) {
		n.Run(serverAddr)
	}(me.Addr)

}

func NewDMetrics(option *Options) *DMetrics {
	glog.V(5).Infoln(` New a metrics... `)
	me := new(DMetrics)
	me.Metrics = make(map[string]interface{})
	me.Path = option.Path
	me.Addr = option.Addr

	if option.Switcher == 1 {
		InitDMetrics(me)
	}
	return me
}

// register Metrics
func (me *DMetrics) RegMetric(name string, metric interface{}) interface{} {
	glog.V(5).Infoln(name, ` Regist metrics... `)
	me.mutex.Lock()
	defer me.mutex.Unlock()
	if _, ok := me.Metrics[name]; ok {
		//已经存在，命名冲突啦
		glog.V(5).Infoln(name, ` RegMetric error: name comflict... `)
	}
	me.Metrics[name] = metric
	return metric
}

// register route
func (me *DMetrics) RegHandler(mux *mux.Router) {
	mux.HandleFunc("/api/metrics/metrics/{action}", me.MetricsResponse)
}

// Metrics数据解析响应
func (me *DMetrics) MetricsResponse(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	metricName := vars["action"]
	glog.V(5).Infoln(metricName, ` action response... `)

	res, err := me.metrics2string(metricName)
	if err != nil {
		glog.V(5).Infoln(err)
	}

	w.Write([]byte(res))
}

func (me *DMetrics) metrics2string(metricName string) (string, error) {
	var res string
	switch value := me.Metrics[metricName]; vtype := me.Metrics[metricName].(type) {
	case metricslib.Counter:
		res = metrics2json(reflect.TypeOf(value.(metricslib.Counter)), reflect.ValueOf(value.(metricslib.Counter)))
	case metricslib.Gauge:
		res = metrics2json(reflect.TypeOf(value.(metricslib.Gauge)), reflect.ValueOf(value.(metricslib.Gauge)))
	case metricslib.EWMA:
		res = metrics2json(reflect.TypeOf(value.(metricslib.EWMA)), reflect.ValueOf(value.(metricslib.EWMA)))
	case metricslib.Mstring:
		res = metrics2json(reflect.TypeOf(value.(metricslib.Mstring)), reflect.ValueOf(value.(metricslib.Mstring)))
	case metricslib.Timer:
		res = metrics2json(reflect.TypeOf(value.(metricslib.Timer)), reflect.ValueOf(value.(metricslib.Timer)))
	case metricslib.Sample:
		res = metrics2json(reflect.TypeOf(value.(metricslib.Sample)), reflect.ValueOf(value.(metricslib.Sample)))
	case metricslib.Histogram:
		res = metrics2json(reflect.TypeOf(value.(metricslib.Histogram)), reflect.ValueOf(value.(metricslib.Histogram)))
	case metricslib.Meter:
		res = metrics2json(reflect.TypeOf(value.(metricslib.Meter)), reflect.ValueOf(value.(metricslib.Meter)))
	case metricslib.Flow:
		res = metrics2json(reflect.TypeOf(value.(metricslib.Flow)), reflect.ValueOf(value.(metricslib.Flow)))
	default:
		{
			strs, err := json.Marshal(struct {
				Error string `json:"error"`
			}{"Unknown"})
			if err != nil {
				return "", err
			}
			glog.V(5).Infoln(vtype)
			res = string(strs)
		}
	}
	return res, nil
}

func metrics2json(t reflect.Type, v reflect.Value) string {
	data := make(map[string]interface{})
	for i := 0; i < t.NumMethod(); i++ {
		if t.Method(i).Type.NumIn() == 1 && t.Method(i).Type.NumOut() == 1 {
			switch v.Method(i).Call([]reflect.Value{})[0].Interface().(type) {
			case int64:
				data[t.Method(i).Name] = strconv.FormatInt(v.Method(i).Call([]reflect.Value{})[0].Interface().(int64), 10)
			case string:
				data[t.Method(i).Name] = v.Method(i).Call([]reflect.Value{})[0].Interface().(string)
			case []float64:
				{
					var sli []string
					peres := v.Method(i).Call([]reflect.Value{})[0].Interface().([]float64)
					for index := 0; index < len(peres); index++ {
						sli = append(sli, strconv.FormatFloat(peres[index], 'G', 10, 64))
					}
					data[t.Method(i).Name] = sli
				}
			case []int64:
				{
					var sli []string
					peres := v.Method(i).Call([]reflect.Value{})[0].Interface().([]int64)
					for index := 0; index < len(peres); index++ {
						sli = append(sli, strconv.FormatInt(peres[index], 10))
					}
					data[t.Method(i).Name] = sli
				}
			case []([]byte):
				{
					var sli []string
					peres := v.Method(i).Call([]reflect.Value{})[0].Interface().([][]byte)
					for index := 0; index < len(peres); index++ {
						sli = append(sli, string(peres[index][0:len(peres[index])]))
					}
					data[t.Method(i).Name] = sli
				}
			case float64:
				data[t.Method(i).Name] = strconv.FormatFloat(v.Method(i).Call([]reflect.Value{})[0].Interface().(float64), 'G', 10, 64)
			}
		}
	}

	strs, err := json.Marshal(data)
	if err != nil {
		glog.V(5).Infoln(err)
	}
	return string(strs)
}
