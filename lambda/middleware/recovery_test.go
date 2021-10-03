package middleware_test

import (
	"errors"
	"github.com/drpsychick/go-alexa-lambda"
	"github.com/drpsychick/go-alexa-lambda-template/lambda/middleware"
	"github.com/hamba/logger/v2"
	"os"
	"testing"
)

func TestWithRecovery(t *testing.T) {
	h := middleware.WithRecovery(
		alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
			panic("panic")
		}),
		logger.New(os.Stdout, logger.ConsoleFormat(), logger.Info),
	)

	bdr := &alexa.ResponseBuilder{}
	req := &alexa.RequestEnvelope{}

	defer func() {
		if err := recover(); err != nil {
			t.Fatal("Expected the panic to be handled.")
		}
	}()

	h.Serve(bdr, req)
}

func TestWithRecovery_Error(t *testing.T) {
	h := middleware.WithRecovery(
		alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
			panic(errors.New("panic"))
		}),
		logger.New(os.Stdout, logger.ConsoleFormat(), logger.Info),
	)

	bdr := &alexa.ResponseBuilder{}
	req := &alexa.RequestEnvelope{}

	defer func() {
		if err := recover(); err != nil {
			t.Fatal("Expected the panic to be handled.")
		}
	}()

	h.Serve(bdr, req)
}
