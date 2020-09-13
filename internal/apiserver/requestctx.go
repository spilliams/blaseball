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

func loggerFromRequest(r *http.Request) (*logrus.Entry, error) {
	logger, ok := r.Context().Value(loggerKey).(*logrus.Entry)

	if !ok {
		// make one
		u, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}

		logger = logrus.WithFields(logrus.Fields{
			"method":      r.Method,
			"path":        r.URL.EscapedPath(),
			"requestUUID": u.String(),
		})
	}

	return logger, nil
}

func requestWithLogger(r *http.Request, logger *logrus.Entry) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), loggerKey, logger))
}
