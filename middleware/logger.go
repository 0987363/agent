package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/0987363/agent/models"

	"net"
	"net/url"
	"time"
)

var logConn net.Conn
var logLevel string

func LoggerInit() *logrus.Logger {
	logger := logrus.New()

	logger.Formatter = &logrus.TextFormatter{ForceColors: true, FullTimestamp: true, TimestampFormat: time.RFC3339Nano}

	return logger
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := LoggerInit()

		u, err := url.QueryUnescape(c.Request.URL.String())
		if err != nil {
			u = c.Request.URL.String()
		}
		start := time.Now()

		c.Set(models.MiddwareKeyLogger, log)
		c.Next()

		if c.Request.Method == "OPTIONS" {
			return
		}
		log.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"user_agent": c.Request.UserAgent(),
			"url":        u,
			"remote":     c.ClientIP(),
			"status":     c.Writer.Status(),
			"spent":      int(time.Now().Sub(start) / time.Millisecond),
		}).Infof("Responded %03d in %s", c.Writer.Status(), time.Now().Sub(start))
	}
}

func GetLogger(c *gin.Context) *logrus.Logger {
	if logger, ok := c.Get(models.MiddwareKeyLogger); ok {
		return logger.(*logrus.Logger)
	}

	return nil
}
