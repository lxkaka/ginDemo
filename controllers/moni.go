package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//HTTPReqDuration metric:http_request_duration_seconds
	HTTPReqDuration *prometheus.HistogramVec
	//HTTPReqTotal metric:http_request_total
	HTTPReqTotal *prometheus.CounterVec
	// TaskRunning metric:task_running
	TaskRunning *prometheus.GaugeVec
)

func init() {
	// 监控接口请求耗时
	// 指标类型是 Histogram
	HTTPReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "http request latencies in seconds",
		Buckets: nil,
	}, []string{"method", "path"})
	// "method"、"path" 是 label

	// 监控接口请求次数
	// 指标类型是 Counter
	HTTPReqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "total number of http requests",
	}, []string{"method", "path", "status"})
	// "method"、"path"、"status" 是 label

	// 监控当前在执行的 task 数量
	// 监控类型是 Gauge
	TaskRunning = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "task_running",
		Help: "current count  of running task",
	}, []string{"type", "state"})
	// "type"、"state" 是 label

	prometheus.MustRegister(
		HTTPReqDuration,
		HTTPReqTotal,
		TaskRunning,
	)
}

func Monitor(c *gin.Context) {
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
}
