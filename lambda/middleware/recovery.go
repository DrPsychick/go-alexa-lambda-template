// Package middleware contains middlewares for lambda.
package middleware

import (
	"fmt"

	"github.com/drpsychick/go-alexa-lambda"
	"github.com/hamba/logger/v2"
)

// Recovery is a middleware that will recover from panics and logs the error.
type Recovery struct {
	handler alexa.Handler
	l       *logger.Logger
}

// WithRecovery recovers from panics and log the error.
func WithRecovery(h alexa.Handler, log *logger.Logger) alexa.Handler {
	return &Recovery{
		handler: h,
		l:       log,
	}
}

// Serve serves the request.
func (m Recovery) Serve(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
	defer func() {
		if v := recover(); v != nil {
			m.l.Error(fmt.Sprintf("%+v", v))
		}
	}()

	m.handler.Serve(b, r)
}
