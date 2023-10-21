package main

import (
	"log"
	"os"

	"github.com/hamba/cmd/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
)

var version = "v0.0.1"

var commands = []*cli.Command{
	{
		Name:   "server",
		Usage:  "Run the lambda server",
		Action: runServer,
		Flags:  cmd.Flags{}.Merge(cmd.LogFlags, cmd.StatsFlags, serverFlags),
	},
	{
		Name:   "lambda",
		Usage:  "Run the lambda server",
		Action: runLambda,
		Flags: cmd.Flags{
			&cli.IntFlag{
				Name:    "lambda.port",
				Usage:   "Port on which lambda will listen",
				EnvVars: []string{"_LAMBDA_SERVER_PORT"},
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
				EnvVars: []string{"ALEXA_MAKE_SKILL"},
			},
			&cli.BoolFlag{
				Name:    "models",
				Usage:   "Generate Alexa interaction model JSON files",
				EnvVars: []string{"ALEXA_MAKE_MODELS"},
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
		EnvVars: []string{"PORT"},
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "My Demo Skill"
	app.Usage = "Build skill and run lambda to answer the Skills requests"
	app.Version = version
	app.Commands = commands
	// need to be set for default Action
	app.Flags = sharedFlags.Merge(cmd.LogFlags, cmd.StatsFlags, serverFlags)
	app.Action = runLambda

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
