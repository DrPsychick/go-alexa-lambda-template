package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hamba/cmd/v3/observe"
	"github.com/urfave/cli/v3"
)

func runMake(ctx context.Context, cmd *cli.Command) error {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	obsvr, err := observe.New(ctx, cmd, "tocli", &observe.Options{
		StatsRuntime: false,
	})
	if err != nil {
		return fmt.Errorf("failed to create observer: %w", err)
	}
	defer obsvr.Close()

	app, err := newApplication(obsvr)
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
	if cmd.Bool("skill") {
		if err := os.MkdirAll("./alexa", 0o775); err != nil {
			app.Logger().Error("Could not create './alexa' directory!")
			return err
		}
		s, err := sk.Build()
		if err != nil {
			return err
		}
		res, _ := json.MarshalIndent(s, "", "  ")
		if err := os.WriteFile("./alexa/skill.json", res, 0o644); err != nil {
			return err
		}
	}

	if cmd.Bool("models") {
		if err := os.MkdirAll("./alexa/interactionModels/custom", 0o755); err != nil {
			app.Logger().Error("Could not create './alexa/interactionModels/custom' directory!")
			return err
		}
		for l, m := range ms {
			filename := "./alexa/interactionModels/custom/" + l + ".json"

			res, _ := json.MarshalIndent(m, "", "  ")
			if err := os.WriteFile(filename, res, 0o644); err != nil {
				app.Logger().Error(err.Error())
				return err
			}
		}
	}

	return nil
}
