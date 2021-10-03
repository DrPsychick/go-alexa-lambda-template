// Package lambda defines how to handle requests and returns proper responses.
package lambda

import (
	"fmt"
	mydemoskill "github.com/drpsychick/go-alexa-lambda-template"

	"github.com/drpsychick/go-alexa-lambda"
	"github.com/drpsychick/go-alexa-lambda-template/loca"
	"github.com/drpsychick/go-alexa-lambda/l10n"
	"github.com/drpsychick/go-alexa-lambda/skill"
	"github.com/hamba/logger/v2"
	"github.com/hamba/statter/v2"
)

// Application defines the interface used of the app.
type Application interface {
	Logger() *logger.Logger
	Statter() *statter.Statter

	Launch(l l10n.LocaleInstance, opts ...mydemoskill.ResponseFunc) (alexa.Response, error)
	Help(l l10n.LocaleInstance, opts ...mydemoskill.ResponseFunc) (alexa.Response, error)
	Stop(l l10n.LocaleInstance, opts ...mydemoskill.ResponseFunc) (alexa.Response, error)
	Cancel(l l10n.LocaleInstance, opts ...mydemoskill.ResponseFunc) (alexa.Response, error)
	DoSomething(
		l l10n.LocaleInstance, server string, opts ...mydemoskill.ResponseFunc,
	) (alexa.Response, error)
}

// NewMux returns a new handler for defined intents.
func NewMux(app Application, sb *skill.SkillBuilder) alexa.Handler {
	mux := alexa.NewServerMux(app.Logger())
	sb.WithModel()

	mux.HandleRequestTypeFunc(alexa.TypeLaunchRequest, handleLaunch(app))
	mux.HandleRequestTypeFunc(alexa.TypeSessionEndedRequest, handleStop(app, sb))

	// register intents
	mux.HandleIntent(alexa.HelpIntent, handleHelp(app, sb))
	mux.HandleIntent(alexa.CancelIntent, handleCancel(app, sb))
	mux.HandleIntent(alexa.StopIntent, handleStop(app, sb))

	mux.HandleIntent(loca.DoSomething, handleDoSomething(app, sb))

	// this is called last, so it can extract all intents from the model
	mux.HandleRequestTypeFunc(alexa.TypeCanFulfillIntentRequest, handleCanFulfillIntent(app, sb))

	return mux
}

func validateSlot(r *alexa.RequestEnvelope, slot string) (string, error) {
	s, err := r.Slot(slot)
	if err != nil {
		return "", err
	}
	if _, err := s.FirstAuthorityWithMatch(); err != nil {
		return "", err
	}
	return r.SlotValue(slot), nil
}

func doSomething(app Application, loc l10n.LocaleInstance, b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) error {
	task, err := validateSlot(r, loca.TypeTaskName)
	if err != nil {
		return err
	}

	// add person if given in context
	responseFuncs := []mydemoskill.ResponseFunc{}
	per, err := r.ContextPerson()
	if err == nil {
		responseFuncs = append(responseFuncs, mydemoskill.WithUser(per.PersonID))
	}

	resp, err := app.DoSomething(loc, task, responseFuncs...)
	if err != nil {
		return err
	}
	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

func handleDoSomething(app Application, sb *skill.SkillBuilder) alexa.Handler {
	m := sb.Model().WithIntent(loca.DoSomething).
		WithType(loca.TypeTask)

	m.Intent(loca.DoSomething).
		WithDelegation(skill.DelegationSkillResponse).
		WithSlot(loca.TypeTaskName, loca.TypeTask)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			app.Logger().Error(err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := doSomething(app, loc, b, r); err != nil {
			app.Logger().Error("could not handle doSomething: " + err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

func handleCanFulfillIntent(app Application, sb *skill.SkillBuilder) alexa.HandlerFunc {
	intents := sb.Model().Intents()

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		intent, err := r.Intent()
		if err != nil {
			b.WithCanFulfillIntent(&alexa.CanFulfillIntent{
				CanFulfill: "NO",
			})
			return
		}

		for _, i := range intents {
			if intent.Name == i {
				b.WithCanFulfillIntent(&alexa.CanFulfillIntent{
					CanFulfill: "YES",
				})
				return
			}
		}
		app.Logger().Info(fmt.Sprintf("could not fulfill intent: %s", intent.Name))

		b.WithCanFulfillIntent(&alexa.CanFulfillIntent{
			CanFulfill: "NO",
		})
	})
}

func launch(app Application, loc l10n.LocaleInstance, b *alexa.ResponseBuilder) error {
	resp, err := app.Launch(loc)
	if err != nil {
		return err
	}
	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

func handleLaunch(app Application) alexa.HandlerFunc {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			app.Logger().Error(err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := launch(app, loc, b); err != nil {
			app.Logger().Error("could not handle Stop: " + err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

func cancel(app Application, loc l10n.LocaleInstance, b *alexa.ResponseBuilder) error {
	resp, err := app.Cancel(loc)
	if err != nil {
		return err
	}
	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

func handleCancel(app Application, sb *skill.SkillBuilder) alexa.HandlerFunc { //nolint:dupl
	sb.Model().WithIntent(alexa.CancelIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			app.Logger().Error(err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := cancel(app, loc, b); err != nil {
			app.Logger().Error("could not handle Stop: " + err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

func stop(app Application, loc l10n.LocaleInstance, b *alexa.ResponseBuilder) error {
	resp, err := app.Stop(loc)
	if err != nil {
		return err
	}
	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

func handleStop(app Application, sb *skill.SkillBuilder) alexa.HandlerFunc { //nolint:dupl
	sb.Model().WithIntent(alexa.StopIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			app.Logger().Error(err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := stop(app, loc, b); err != nil {
			app.Logger().Error("could not handle Stop: " + err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

func help(app Application, loc l10n.LocaleInstance, b *alexa.ResponseBuilder) error {
	resp, err := app.Help(loc)
	if err != nil {
		return err
	}
	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

// handleHelp calls the app help method, it does not close the session.
func handleHelp(app Application, sb *skill.SkillBuilder) alexa.Handler { //nolint:dupl
	sb.Model().WithIntent(alexa.HelpIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			app.Logger().Error(err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := help(app, loc, b); err != nil {
			app.Logger().Error("could not handle Stop: " + err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

// DefaultError is a generic error.
type DefaultError struct {
	Locale l10n.LocaleInstance
}

// Error returns a string.
func (m DefaultError) Error() string {
	return "an error occurred"
}

// Response returns a default error response.
func (m DefaultError) Response(loc l10n.LocaleInstance) alexa.Response {
	return alexa.Response{
		Title:  loc.GetAny(l10n.KeyErrorTitle),
		Text:   loc.GetAny(l10n.KeyErrorText),
		Speech: loc.GetAny(l10n.KeyErrorSSML),
		End:    true,
	}
}
