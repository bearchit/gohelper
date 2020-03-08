package kithelper

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/go-kit/kit/transport"
)

type logrusErrorHandler struct {
	logger logrus.FieldLogger
}

func NewLogrusErrorHandler(logger logrus.FieldLogger) transport.ErrorHandler {
	return &logrusErrorHandler{
		logger: logger,
	}
}

func (h logrusErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Error(err)
}
