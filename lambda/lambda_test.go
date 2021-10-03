package lambda_test

import (
	"github.com/drpsychick/go-alexa-lambda"
	mydemoskill "github.com/drpsychick/go-alexa-lambda-template"
	"github.com/drpsychick/go-alexa-lambda-template/lambda"
	"github.com/drpsychick/go-alexa-lambda-template/lambda/middleware"
	"github.com/drpsychick/go-alexa-lambda-template/loca"
	"github.com/drpsychick/go-alexa-lambda/l10n"
	"github.com/drpsychick/go-alexa-lambda/skill"
	"github.com/drpsychick/go-alexa-lambda/ssml"
	"github.com/hamba/logger/v2"
	"github.com/hamba/statter/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func initLocaleRegistry(t *testing.T) {
	loca.Registry = l10n.NewRegistry()
	err := loca.Registry.Register(&l10n.Locale{Name: "en-US", TextSnippets: l10n.Snippets{}})
	assert.NoError(t, err)
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	loc.Set(l10n.KeyErrorTitle, []string{"error"})
	loc.Set(l10n.KeyErrorText, []string{"An error occurred: %s"})
	loc.Set(l10n.KeyErrorSSML, []string{ssml.Speak("An error occurred.")})
	loc.Set(l10n.KeyErrorLocaleNotFoundTitle, []string{"error"})
	loc.Set(l10n.KeyErrorLocaleNotFoundText, []string{"Locale '%s' not found!"})
	loc.Set(l10n.KeyErrorLocaleNotFoundSSML, []string{ssml.Speak("Locale '%s' not found!")})
	loc.Set(l10n.KeyErrorNoTranslationTitle, []string{"error"})
	loc.Set(l10n.KeyErrorNoTranslationText, []string{"Key '%s' not found!"})
	loc.Set(l10n.KeyErrorNoTranslationSSML, []string{ssml.Speak("Key '%s' not found!")})
	loc.Set(l10n.KeyErrorTranslationTitle, []string{"Translation error"})
	loc.Set(l10n.KeyErrorTranslationText, []string{"An error occurred in translation. The developer is informed."})
	loc.Set(l10n.KeyErrorTranslationSSML, []string{"<speak>An error occurred during translation. The developer is informed.<speak>"})
}

func TestLambda_HandleLaunch(t *testing.T) {
	initLocaleRegistry(t)

	app := mydemoskill.NewApplication(logger.New(os.Stdout, logger.ConsoleFormat(), logger.Info), statter.New(nil, time.Second))
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	assert.NotEmpty(t, loc)

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeLaunchRequest,
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app.Statter())

	m.Serve(b, r)
	resp := b.Build()

	// locale not found
	assert.NotEmpty(t, resp)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	loc.ResetErrors()

	// now with locale
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Contains(t, loc.GetErrors()[0].Error(), l10n.KeyLaunchTitle)
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	loc.ResetErrors()

	// now with loca
	loc.Set(l10n.KeyLaunchTitle, []string{"Start"})
	loc.Set(l10n.KeyLaunchText, []string{"Und los..."})
	loc.Set(l10n.KeyLaunchSSML, []string{"foo", "bar"})
	assert.NotEmpty(t, loc.Get(l10n.KeyLaunchTitle))
	assert.NotEmpty(t, loc.GetAny(l10n.KeyLaunchText))
	assert.NotEmpty(t, loc.GetAny(l10n.KeyLaunchSSML))

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(l10n.KeyLaunchTitle), resp.Response.Card.Title)
}

func TestLambda_HandleEnd(t *testing.T) {
	initLocaleRegistry(t)

	app := mydemoskill.NewApplication(logger.New(os.Stdout, logger.ConsoleFormat(), logger.Info), statter.New(nil, time.Second))

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeSessionEndedRequest,
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app.Statter())
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	loc.ResetErrors()

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Contains(t, loc.GetErrors()[0].Error(), l10n.KeyStopTitle)
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	loc.ResetErrors()

	// with translations
	loc.Set(l10n.KeyStopTitle, []string{"Stop"})
	loc.Set(l10n.KeyStopText, []string{"Alright, it's over now."})
	loc.Set(l10n.KeyStopSSML, []string{"Over and out"})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(l10n.KeyStopTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(l10n.KeyStopText), resp.Response.Card.Content)
}

func TestLambda_HandleHelp(t *testing.T) {
	initLocaleRegistry(t)

	app := mydemoskill.NewApplication(logger.New(os.Stdout, logger.ConsoleFormat(), logger.Info), statter.New(nil, time.Second))

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: alexa.HelpIntent,
			},
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app.Statter())
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	// from here:
	// https://github.com/DrPsychick/go-alexa-lambda/blob/3d9d0de7ef97a361766267a23ec5b06ca1b54862/application.go#L60
	assert.Equal(t, "Error", resp.Response.Card.Title)
	loc.ResetErrors()

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Contains(t, loc.GetErrors()[0].Error(), l10n.KeyHelpTitle)
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	loc.ResetErrors()

	// with translations
	loc.Set(l10n.KeyHelpTitle, []string{"Help"})
	loc.Set(l10n.KeyHelpText, []string{"I'd love to help you"})
	loc.Set(l10n.KeyHelpSSML, []string{"Help"})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(l10n.KeyHelpTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(l10n.KeyHelpText), resp.Response.Card.Content)
}

func TestLambda_HandleStop(t *testing.T) {
	initLocaleRegistry(t)

	app := mydemoskill.NewApplication(logger.New(os.Stdout, logger.ConsoleFormat(), logger.Info), statter.New(nil, time.Second))

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: alexa.StopIntent,
			},
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app.Statter())
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	loc.ResetErrors()

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Contains(t, loc.GetErrors()[0].Error(), l10n.KeyStopTitle)
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	loc.ResetErrors()

	// with translations

	loc.Set(l10n.KeyStopTitle, []string{"Stop"})
	loc.Set(l10n.KeyStopText, []string{"Alright, it's over now."})
	loc.Set(l10n.KeyStopSSML, []string{"Over and out"})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(l10n.KeyStopTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(l10n.KeyStopText), resp.Response.Card.Content)
}
