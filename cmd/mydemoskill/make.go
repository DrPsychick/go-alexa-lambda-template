package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
)

func runMake(c *cli.Context) error {
	app, err := newApplication(c)
	if err != nil {
		return err
	}

	// build skill and models
	sk := newSkill()

	// lambda injects supported intents, slots, types
	newLambda(app, sk)

	ms, err := createSkillModels(sk)
	if err != nil {
		return err
	}

	// build and write JSON files
	if c.Bool("skill") {
		if err := os.MkdirAll("./alexa", 0o775); err != nil {
			app.Logger().Error("Could not create './alexa' directory!")
			return err
		}
		s, err := sk.Build()
		if err != nil {
			return err
		}
		res, _ := json.MarshalIndent(s, "", "  ")
		if err := ioutil.WriteFile("./alexa/skill.json", res, 0o644); err != nil {
			return err
		}
	}

	if c.Bool("models") {
		if err := os.MkdirAll("./alexa/interactionModels/custom", 0o755); err != nil {
			app.Logger().Error("Could not create './alexa/interactionModels/custom' directory!")
			return err
		}
		for l, m := range ms {
			filename := "./alexa/interactionModels/custom/" + l + ".json"

			res, _ := json.MarshalIndent(m, "", "  ")
			if err := ioutil.WriteFile(filename, res, 0o644); err != nil {
				app.Logger().Error(err.Error())
				return err
			}
		}
	}

	return nil
}
