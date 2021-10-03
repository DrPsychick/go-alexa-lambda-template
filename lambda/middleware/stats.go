package middleware

import (
	"strings"

	"github.com/drpsychick/go-alexa-lambda"
	"github.com/hamba/statter/v2"
	"github.com/hamba/timex/mono"
)

// WithRequestStats adds counter and timing stats to intent requests.
func WithRequestStats(h alexa.Handler, stat *statter.Statter) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		tags := []statter.Tag{{"locale", r.RequestLocale()}}

		if r.IsIntentRequest() {
			tags = append(tags, statter.Tag{"intent", r.IntentName()})
		}
		for _, s := range r.Slots() {
			if s.Value != "" {
				tags = append(tags, statter.Tag{strings.ToLower(s.Name), s.Value})
			}
		}

		stat.Counter("request.start", tags...).Inc(1)
		start := mono.Now()
		h.Serve(b, r)

		d := mono.Since(start)

		stat.Timing("request.time", tags...).Observe(d)
		stat.Counter("request.complete", tags...).Inc(1)
	})
}
