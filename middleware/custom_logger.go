package middleware

import (
	"encoding/json"
	"strings"
	// "time"

	"github.com/gin-gonic/gin"
)

type logEntry struct {
	Method  string            `json:"method"`
	Path    string            `json:"path"`
	Status  int               `json:"status"`
	// Latency string            `json:"latency"`
	Headers map[string]string `json:"headers"` // <- declarado por último
}

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// start := time.Now()
		c.Next()
		// duration := time.Since(start)

		path := c.Request.URL.Path
		status := c.Writer.Status()

		if strings.HasPrefix(path, "/static/") && status < 400 {
			return
		}

		// Build headers
		headers := make(map[string]string)
		for k, v := range c.Request.Header {
			headers[k] = strings.Join(v, ", ")
		}

		// Struct enforces field order
		entry := logEntry{
			Method:  c.Request.Method,
			Path:    path,
			Status:  status,
			// Latency: duration.String(),
			Headers: headers, // último no struct → último no JSON
		}

		jsonLog, _ := json.Marshal(entry)
		LoggerInfo("request" + string(jsonLog))
	}
}
