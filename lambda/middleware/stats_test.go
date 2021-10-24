package middleware_test

import (
	"github.com/drpsychick/go-alexa-lambda"
	"github.com/drpsychick/go-alexa-lambda-template/lambda/middleware"
	"github.com/hamba/statter/v2"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestWithRequestStats(t *testing.T) {
	tags := [][2]string{{"intent", "test-intent"}, {"locale", "en-US"}, {"test-slot", "slot-value"}}
	r := new(mockSimpleReporter)
	r.On("Counter", "request.start", int64(1), tags)
	r.On("Timing", "request.time", tags)
	r.On("Counter", "request.complete", int64(1), tags)
	stats := statter.New(r, time.Second)
	m := middleware.WithRequestStats(alexa.HandlerFunc(
		func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {}),
		stats,
	)

	bdr := &alexa.ResponseBuilder{}
	req := &alexa.RequestEnvelope{
		Request: &alexa.Request{
			Type: alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: "test-intent",
				Slots: map[string]*alexa.Slot{
					"test-slot": {
						Name:  "test-slot",
						Value: "slot-value",
					},
				},
			},
			Locale: "en-US",
		},
	}

	m.Serve(bdr, req)
	stats.Close()
	r.AssertExpectations(t)
}

func TestWithRequestStats_NonIntentRequests(t *testing.T) {
	tags := [][2]string{{"locale", "en-US"}}
	r := new(mockSimpleReporter)
	r.On("Counter", "request.start", int64(1), tags)
	r.On("Timing", "request.time", tags)
	r.On("Counter", "request.complete", int64(1), tags)
	stats := statter.New(r, time.Second)
	m := middleware.WithRequestStats(alexa.HandlerFunc(
		func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {}),
		stats,
	)

	bdr := &alexa.ResponseBuilder{}
	req := &alexa.RequestEnvelope{
		Request: &alexa.Request{
			Type:   alexa.TypeLaunchRequest,
			Locale: "en-US",
		},
	}

	m.Serve(bdr, req)
	stats.Close()
	r.AssertExpectations(t)
}

type mockSimpleReporter struct {
	mock.Mock
}

func (r *mockSimpleReporter) Counter(name string, v int64, tags [][2]string) {
	_ = r.Called(name, v, tags)
}

func (r *mockSimpleReporter) Gauge(name string, v float64, tags [][2]string) {
	_ = r.Called(name, v, tags)
}

func (r *mockSimpleReporter) Timing(name string, tags [][2]string) func(v time.Duration) {
	_ = r.Called(name, tags)
	return func(v time.Duration) {}
}

func (r *mockSimpleReporter) Close() error {
	return nil
}
