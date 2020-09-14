package apiserver

import (
	"context"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type ctxKey string

const (
	loggerKey ctxKey = "logger"
)

func loggerFromRequest(r *http.Request) *logrus.Entry {
	logger, ok := r.Context().Value(loggerKey).(*logrus.Entry)

	if !ok {
		// make one
		u, err := uuid.NewV4()
		uString := "unknown"
		if u != [16]byte{} {
			uString = u.String()
		}

		logger = logrus.WithFields(logrus.Fields{
			"method":      r.Method,
			"path":        r.URL.EscapedPath(),
			"requestUUID": uString,
		})

		if err != nil {
			logger.Warnf("could not generate uuid for logger: %v", err)
		}
	}

	return logger
}

func requestWithLogger(r *http.Request, logger *logrus.Entry) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), loggerKey, logger))
}
