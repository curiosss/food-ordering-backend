package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	logger *logrus.Logger
}

func NewLogMiddleware(logger *logrus.Logger) *LogMiddleware {
	return &LogMiddleware{logger}
}

func (l *LogMiddleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// l.logger.Infof("%s -> from ip: %s url: %s", r.Method, r.Host, r.URL)
		l.logger.WithFields(logrus.Fields{
			"host": r.Host,
			"m":    r.Method,
			"url":  r.URL,
			"z":    r.Header,
		}).Info()
		// l.logger.Info("")
		next.ServeHTTP(w, r)
	})
}
