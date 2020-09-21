package middlewares

import (
	"ginDemo/controllers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"strconv"
	"time"
)

func shuffle(s []string) string {
	r := rand.Intn(2)
	return s[r]
}

//Metric metric middleware
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		tBegin := time.Now()
		c.Next()

		duration := float64(time.Since(tBegin)) / float64(time.Second)

		path := c.Request.URL.Path

		// 请求数加1
		controllers.HTTPReqTotal.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   path,
			"status": strconv.Itoa(c.Writer.Status()),
		}).Inc()

		//  记录本次请求处理时间
		controllers.HTTPReqDuration.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   path,
		}).Observe(duration)

		controllers.TaskRunning.With(prometheus.Labels{
			"type":  shuffle([]string{"video", "audio"}),
			"state": shuffle([]string{"process", "queue"}),
		}).Inc()

		controllers.TaskRunning.With(prometheus.Labels{
			"type":  shuffle([]string{"video", "audio"}),
			"state": shuffle([]string{"process", "queue"}),
		}).Dec()
	}
}
