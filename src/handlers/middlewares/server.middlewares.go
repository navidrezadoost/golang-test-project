package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func ServerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logger middleware
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("[%s] %s %s %s %d %s\n",
			start.Format("2006/01/02 - 15:04:05"),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)

		// Recovery middleware
		if err := c.Errors.Last(); err != nil {
			log.Printf("Panic error: %s\n", err.Error())
		}
	}
}
