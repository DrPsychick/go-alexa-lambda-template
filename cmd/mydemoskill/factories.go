package main

import (
	mydemoskill "github.com/drpsychick/go-alexa-lambda-template"
	"github.com/drpsychick/go-alexa-lambda/skill"
	"github.com/hamba/cmd/v3/observe"
)

func newApplication(obsvr *observe.Observer) (*mydemoskill.Application, error) {
	app := mydemoskill.NewApplication(
		obsvr.Log,
		obsvr.Stats,
	)

	return app, nil
}

func newSkill() *skill.SkillBuilder {
	return mydemoskill.NewSkill()
}

func createSkillModels(s *skill.SkillBuilder) (map[string]*skill.Model, error) {
	return mydemoskill.CreateSkillModels(s)
}
