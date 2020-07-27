package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func AccessLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintln("%s - [%s] %s %s %d %s \"%s\"",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			)
	})
}
