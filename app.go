// Package mydemoskill is the core app functionality.
package mydemoskill

import (
	"errors"
	"github.com/drpsychick/go-alexa-lambda-template/loca"
	"github.com/hamba/logger/v2"
	"github.com/hamba/statter/v2"

	alexa "github.com/drpsychick/go-alexa-lambda"
	"github.com/drpsychick/go-alexa-lambda/l10n"
)

const (
	errorUser string = "errorUser"
)

func (a *Application) taskInList(loc l10n.LocaleInstance, task string) bool {
	for _, s := range loc.GetAll(loca.TypeTaskValues) {
		if s == task {
			return true
		}
	}
	return false
}

// DoSomething triggers the start of a server and returns the result.
func (a *Application) DoSomething(loc l10n.LocaleInstance, task string, opts ...ResponseFunc, //nolint:dupl
) (alexa.Response, error) {
	// run all ResponseFuncs
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	// check input: does the task exist?
	if !a.taskInList(loc, task) {
		return a.ElicitServer(loc, task)
	}

	// execute: trigger start of the task

	var tit, msg, msgSSML string
	if cfg.User != "" {
		// personalized response
		tit = loc.GetAny(loca.DoSomethingTitle, cfg.User)
		msg = loc.GetAny(loca.DoSomethingUserText, cfg.User, task)
		msgSSML = loc.GetAny(loca.DoSomethingUserSSML, cfg.User, task)
	} else {
		tit = loc.GetAny(loca.DoSomethingTitle)
		msg = loc.GetAny(loca.DoSomethingText, task)
		msgSSML = loc.GetAny(loca.DoSomethingSSML, task)
	}

	if cfg.User == errorUser {
		return alexa.Response{}, ErrUnknown
	}

	return alexa.Response{
		Title:  tit,
		Text:   msg,
		Speech: msgSSML,
		End:    true,
	}, nil
}

// ElicitServer reprompts for a valid server name.
func (a *Application) ElicitServer(loc l10n.LocaleInstance, server string) (alexa.Response, error) {
	resp := alexa.Response{
		Title:    loc.GetAny(loca.SlotTaskElicitTitle),
		Reprompt: false,
		End:      false,
	}

	if server != "" {
		resp.Text = loc.GetAny(loca.SlotWrongTaskElicitText, server)
		resp.Speech = loc.GetAny(loca.SlotWrongTaskElicitSSML, server)
		return resp, nil
	}
	resp.Text = loc.GetAny(loca.SlotTaskElicitText)
	resp.Speech = loc.GetAny(loca.SlotTaskElicitSSML)
	return resp, nil
}

// Launch starts a skill session.
func (a *Application) Launch(loc l10n.LocaleInstance, opts ...ResponseFunc) (alexa.Response, error) {
	// run all ResponseFuncs
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	// check input: does the server exist?

	// execute: trigger start of the server

	var tit, msg, msgSSML string
	// if cfg.User != "" {
	//	// personalized response
	//	tit = loc.GetAny(loca.DoSomethingTitle, cfg.User)
	//	msg = loc.GetAny(loca.DoSomethingText, cfg.User)
	//	msgSSML = loc.GetAny(loca.DoSomethingSSML, cfg.User)
	// } else {
	tit = loc.GetAny(l10n.KeyLaunchTitle)
	msg = loc.GetAny(l10n.KeyLaunchText)
	msgSSML = loc.GetAny(l10n.KeyLaunchSSML)
	// }

	if cfg.User == errorUser {
		return alexa.Response{}, ErrUnknown
	}

	return alexa.Response{
		Title:  tit,
		Text:   msg,
		Speech: msgSSML,
		End:    false,
	}, nil
}

// Help returns a response that explains how to use the skill.
func (a *Application) Help(loc l10n.LocaleInstance, opts ...ResponseFunc) (alexa.Response, error) {
	// run all ResponseFuncs
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	var tit, msg, msgSSML string
	// if cfg.User != "" {
	//	// personalized response
	//	tit = loc.GetAny(loca.DoSomethingTitle, cfg.User)
	//	msg = loc.GetAny(loca.DoSomethingText, cfg.User)
	//	msgSSML = loc.GetAny(loca.DoSomethingSSML, cfg.User)
	// } else {
	tit = loc.GetAny(l10n.KeyHelpTitle)
	msg = loc.GetAny(l10n.KeyHelpText)
	msgSSML = loc.GetAny(l10n.KeyHelpSSML)
	// }

	if cfg.User == errorUser {
		return alexa.Response{}, ErrUnknown
	}

	return alexa.Response{
		Title:  tit,
		Text:   msg,
		Speech: msgSSML,
		End:    false,
	}, nil
}

// Stop ends the skill session.
func (a *Application) Stop(loc l10n.LocaleInstance, opts ...ResponseFunc) (alexa.Response, error) {
	// run all ResponseFuncs
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	var tit, msg, msgSSML string
	// if cfg.User != "" {
	//	// personalized response
	//	tit = loc.GetAny(loca.DoSomethingTitle, cfg.User)
	//	msg = loc.GetAny(loca.DoSomethingText, cfg.User)
	//	msgSSML = loc.GetAny(loca.DoSomethingSSML, cfg.User)
	// } else {
	tit = loc.GetAny(l10n.KeyStopTitle)
	msg = loc.GetAny(l10n.KeyStopText)
	msgSSML = loc.GetAny(l10n.KeyStopSSML)
	// }

	if cfg.User == errorUser {
		return alexa.Response{}, ErrUnknown
	}

	return alexa.Response{
		Title:  tit,
		Text:   msg,
		Speech: msgSSML,
		End:    true,
	}, nil
}

// Cancel cancels the skill session.
func (a *Application) Cancel(loc l10n.LocaleInstance, opts ...ResponseFunc) (alexa.Response, error) {
	// run all ResponseFuncs
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	var tit, msg, msgSSML string
	// if cfg.User != "" {
	//	// personalized response
	//	tit = loc.GetAny(loca.DoSomethingTitle, cfg.User)
	//	msg = loc.GetAny(loca.DoSomethingText, cfg.User)
	//	msgSSML = loc.GetAny(loca.DoSomethingSSML, cfg.User)
	// } else {
	tit = loc.GetAny(l10n.KeyCancelTitle)
	msg = loc.GetAny(l10n.KeyCancelText)
	msgSSML = loc.GetAny(l10n.KeyCancelSSML)
	// }

	if cfg.User == errorUser {
		return alexa.Response{}, ErrUnknown
	}

	return alexa.Response{
		Title:  tit,
		Text:   msg,
		Speech: msgSSML,
		End:    true,
	}, nil
}

// ErrUnknown is the fallback error.
var ErrUnknown = errors.New("something went wrong")

// Application defines the base application.
type Application struct {
	logger  *logger.Logger
	statter *statter.Statter
}

// NewApplication returns an Application with the logger and statter.
func NewApplication(l *logger.Logger, s *statter.Statter) *Application {
	return &Application{
		logger:  l,
		statter: s,
	}
}

// Logger returns the application logger.
func (a *Application) Logger() *logger.Logger {
	return a.logger
}

// Statter returns the application statter.
func (a *Application) Statter() *statter.Statter {
	return a.statter
}

// Config defines additional data that can be provided and used in requests.
type Config struct {
	User string
}

// ResponseFunc defines the function that can optionally be passed to responses.
type ResponseFunc func(cfg *Config)

// WithUser returns a ResponseFunc that sets the user.
func WithUser(user string) ResponseFunc {
	return func(cfg *Config) {
		cfg.User = user
	}
}
