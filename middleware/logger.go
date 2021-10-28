//nolint:exhaustivestruct
package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type Logger struct {
	Log   *logrus.Logger
	start time.Time
}

func NewLogger() *Logger {
	log := logrus.New()

	log.SetReportCaller(false)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})

	return &Logger{
		Log:   log,
		start: time.Now(),
	}
}

func (logger *Logger) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	logger.started(request)

	next(response, request)

	logger.finished(response, request)
}

func (logger *Logger) started(request *http.Request) {
	requestID := request.Header.Get("X-Request-ID")

	logger.Log.WithFields(logrus.Fields{
		"X-Request-ID": requestID,
		"method":       request.Method,
		"request":      request.RequestURI,
		"hostname":     request.Host,
	}).Info("Started handling request")

	logger.start = time.Now()
}

func (logger *Logger) finished(response http.ResponseWriter, request *http.Request) {
	requestID := request.Header.Get("X-Request-ID")

	logger.Log.WithFields(logrus.Fields{
		"X-Request-ID": requestID,
		"method":       request.Method,
		"hostname":     request.Host,
		"request":      request.RequestURI,
		"status":       response.(negroni.ResponseWriter).Status(),
		"text_status":  http.StatusText(response.(negroni.ResponseWriter).Status()),
		"took":         time.Since(logger.start),
		"latency":      time.Since(logger.start).Nanoseconds(),
		"size":         response.(negroni.ResponseWriter).Size(),
	}).Info("Finish handling request")
}
