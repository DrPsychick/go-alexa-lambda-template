package main

import (
	"encoding/json"
	mydemoskill "github.com/drpsychick/go-alexa-lambda-template"
	"github.com/hamba/logger/v2"
	"github.com/hamba/statter/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestMakeSkill(t *testing.T) {
	sb := newSkill()

	s, err := sb.Build()
	assert.NoError(t, err)

	res, err := json.MarshalIndent(s, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
}

func TestMakeModels(t *testing.T) {
	l := logger.New(os.Stdout, logger.ConsoleFormat(), logger.Info)
	s := statter.New(nil, time.Duration(10*time.Second))
	app := mydemoskill.NewApplication(l, s)
	sb := newSkill()
	newLambda(app, sb)

	ms, err := createSkillModels(sb)
	assert.NoError(t, err)

	for _, m := range ms {
		res, err := json.MarshalIndent(m, "", "  ")
		assert.NoError(t, err)
		assert.NotEmpty(t, string(res))
	}
}
