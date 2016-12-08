package metrics

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// metrics-data options
type Options struct {
	Switcher int
	Path     string
	Addr     string
}

// metrics-data struct
type Metrics struct {
	mutex   sync.RWMutex
	Metrics map[string]interface{}
	Path    string
	Addr    string
}

// init metrics-data colloctor
func InitMetrics(me *Metrics) {
	mux := mux.NewRouter()
	me.RegHandler(mux)

	n := negroni.Classic()
	selfDir, _ := filepath.Abs(me.Path)
	n.Use(negroni.NewStatic(http.Dir(selfDir)))
	n.UseHandler(mux)
	log.Println(` metrics strating... `)
	//start http server
	go func(serverAddr string) {
		n.Run(serverAddr)
	}(me.Addr)

}

func NewMetrics(option *Options) *Metrics {
	log.Println(` New a metrics... `)
	me := new(Metrics)
	me.Metrics = make(map[string]interface{})
	me.Path = option.Path
	me.Addr = option.Addr

	if option.Switcher == 1 {
		InitMetrics(me)
	}
	return me
}

// register Metrics
func (me *Metrics) RegMetric(name string, metric interface{}) interface{} {
	log.Println(name, ` Regist metrics... `)
	me.mutex.Lock()
	defer me.mutex.Unlock()
	if _, ok := me.Metrics[name]; ok {
		//已经存在，命名冲突啦
		log.Println(name, ` RegMetric error: name comflict... `)
	}
	me.Metrics[name] = metric
	return metric
}

// register route
func (me *Metrics) RegHandler(mux *mux.Router) {
	mux.HandleFunc("/api/metrics/metrics/{action}", me.MetricsResponse)
}

// Metrics数据解析响应
func (me *Metrics) MetricsResponse(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	metricName := vars["action"]
	log.Println(metricName, ` action response... `)

	res, err := me.Metrics2string(metricName)
	if err != nil {
		log.Println(err)
	}

	w.Write([]byte(res))
}

func (me *Metrics) Metrics2string(metricName string) (string, error) {
	var res string
	switch value := me.Metrics[metricName]; vtype := me.Metrics[metricName].(type) {
	case Counter:
		res = metrics2json(reflect.TypeOf(value.(Counter)), reflect.ValueOf(value.(Counter)))
	case Gauge:
		res = metrics2json(reflect.TypeOf(value.(Gauge)), reflect.ValueOf(value.(Gauge)))
	case EWMA:
		res = metrics2json(reflect.TypeOf(value.(EWMA)), reflect.ValueOf(value.(EWMA)))
	case Mstring:
		res = metrics2json(reflect.TypeOf(value.(Mstring)), reflect.ValueOf(value.(Mstring)))
	case Timer:
		res = metrics2json(reflect.TypeOf(value.(Timer)), reflect.ValueOf(value.(Timer)))
	case Sample:
		res = metrics2json(reflect.TypeOf(value.(Sample)), reflect.ValueOf(value.(Sample)))
	case Histogram:
		res = metrics2json(reflect.TypeOf(value.(Histogram)), reflect.ValueOf(value.(Histogram)))
	case Meter:
		res = metrics2json(reflect.TypeOf(value.(Meter)), reflect.ValueOf(value.(Meter)))
	case Flow:
		res = metrics2json(reflect.TypeOf(value.(Flow)), reflect.ValueOf(value.(Flow)))
	default:
		{
			strs, err := json.Marshal(struct {
				Error string `json:"error"`
			}{"Unknown"})
			if err != nil {
				return "", err
			}
			log.Println(vtype)
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
			case []string:
				{
					var sli []string
					peres := v.Method(i).Call([]reflect.Value{})[0].Interface().([]string)
					for index := 0; index < len(peres); index++ {
						sli = append(sli, peres[index])
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
		log.Println(err)
	}
	return string(strs)
}
