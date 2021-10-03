package main

import (
	mydemoskill "github.com/drpsychick/go-alexa-lambda-template"

	"github.com/drpsychick/go-alexa-lambda/skill"
	"github.com/hamba/cmd/v2"
	"github.com/urfave/cli/v2"
)

func newApplication(c *cli.Context) (*mydemoskill.Application, error) {
	log, err := cmd.NewLogger(c)
	if err != nil {
		return nil, err
	}
	stat, err := cmd.NewStatter(c, log)
	if err != nil {
		return nil, err
	}

	app := mydemoskill.NewApplication(
		log,
		stat,
	)

	return app, nil
}

func newSkill() *skill.SkillBuilder {
	return mydemoskill.NewSkill()
}

func createSkillModels(s *skill.SkillBuilder) (map[string]*skill.Model, error) {
	return mydemoskill.CreateSkillModels(s)
}
