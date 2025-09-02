package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hamba/cmd/v3"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v3"
)

var version = "v0.0.1"

var commands = []*cli.Command{
	{
		Name:   "server",
		Usage:  "Run the lambda server",
		Action: runServer,
		Flags:  cmd.LogFlags.Merge(cmd.StatsFlags).Merge(serverFlags),
	},
	{
		Name:   "lambda",
		Usage:  "Run the lambda server",
		Action: runLambda,
		Flags: cmd.Flags{
			&cli.IntFlag{
				Name:    "lambda.port",
				Usage:   "Port on which lambda will listen",
				Sources: cli.EnvVars("_LAMBDA_SERVER_PORT"),
			},
		}.Merge(cmd.LogFlags, cmd.StatsFlags, serverFlags),
	},
	{
		Name:  "make",
		Usage: "Make Alexa skill files",
		Flags: cmd.Flags{
			&cli.BoolFlag{
				Name:    "skill",
				Usage:   "Generate Alexa skill.json",
				Sources: cli.EnvVars("ALEXA_MAKE_SKILL"),
			},
			&cli.BoolFlag{
				Name:    "models",
				Usage:   "Generate Alexa interaction model JSON files",
				Sources: cli.EnvVars("ALEXA_MAKE_MODELS"),
			},
		}.Merge(cmd.LogFlags, cmd.StatsFlags),
		Action: runMake,
	},
	{
		Name:   "check",
		Usage:  "Check basic skill requirements",
		Flags:  cmd.Flags{}.Merge(cmd.LogFlags, cmd.StatsFlags),
		Action: runCheck,
	},
}

var sharedFlags = cmd.Flags{}
var serverFlags = cmd.Flags{
	&cli.StringFlag{
		Name:    "port",
		Value:   "80",
		Usage:   "Port for HTTP server to listen on",
		Sources: cli.EnvVars("PORT"),
	},
}

func main() {
	app := cli.Command{}
	app.Name = "My Demo Skill"
	app.Usage = "Build skill and run lambda to answer the Skills requests"
	app.Version = version
	app.Commands = commands
	// need to be set for default Action
	app.Flags = sharedFlags.Merge(cmd.LogFlags, cmd.StatsFlags, serverFlags)
	app.Action = runLambda

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(ctx, os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
