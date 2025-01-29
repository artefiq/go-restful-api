package middleware

import (
    "time"

    "github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        duration := time.Since(start)

        // Log request details
        method := c.Request.Method
        path := c.Request.URL.Path
        status := c.Writer.Status()

        c.Writer.Header().Set("X-Response-Time", duration.String())
        c.JSON(status, gin.H{"method": method, "path": path, "duration": duration.String()})
    }
}
