package main

import (
	"errors"
	"fmt"
	mydemoskill "github.com/drpsychick/go-alexa-lambda-template"

	"github.com/drpsychick/go-alexa-lambda"
	"github.com/drpsychick/go-alexa-lambda-template/lambda"
	"github.com/drpsychick/go-alexa-lambda-template/lambda/middleware"
	"github.com/drpsychick/go-alexa-lambda/skill"
	"github.com/hamba/timex/mono"
	"github.com/urfave/cli/v2"
)

func runLambda(c *cli.Context) error {
	start := mono.Now()

	app, err := newApplication(c)
	if err != nil {
		return err
	}
	d := mono.Since(start)
	app.Statter().Timing("Boot").Observe(d)

	sb := newSkill()
	l := newLambda(app, sb)

	ms, err := sb.BuildModels()
	if err != nil {
		return err
	}
	for l, m := range ms {
		app.Logger().Info(fmt.Sprintf("accepting locale '%s' invocation '%s'", l, m.Model.Language.Invocation))
	}

	app.Statter().Timing("Ready").Observe(mono.Since(start))
	if err := alexa.Serve(l); err != nil {
		return err
	}

	return errors.New("Serve() should not have returned")
}

func newLambda(app *mydemoskill.Application, sb *skill.SkillBuilder) alexa.Handler {
	h := lambda.NewMux(app, sb)

	h = middleware.WithRequestStats(h, app.Statter())
	return middleware.WithRecovery(h, app.Logger())
}
