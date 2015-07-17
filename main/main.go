package main

import (
	"os"

	"github.com/codegangsta/cli"
	cliApp "github.com/docker/libcompose/app"
	"github.com/docker/libcompose/command"
	dockerApp "github.com/docker/libcompose/docker/app"
	"github.com/docker/libcompose/version"
)

func main() {
	factory := &dockerApp.ProjectFactory{}

	app := cli.NewApp()
	app.Name = "docker-compose"
	app.Usage = "Fast, isolated environments using Docker."
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Author = "Docker Compose Contributors"
	app.Email = "https://github.com/docker/libcompose"
	app.Before = cliApp.BeforeApp
	app.Flags = append(command.CommonFlags(), dockerApp.DockerClientFlags()...)
	app.Commands = []cli.Command{
		command.CreateCommand(factory),
		command.UpCommand(factory),
		command.StartCommand(factory),
		command.LogsCommand(factory),
		command.RestartCommand(factory),
		command.StopCommand(factory),
		command.ScaleCommand(factory),
		command.RmCommand(factory),
		command.PullCommand(factory),
		command.KillCommand(factory),
	}

	app.Run(os.Args)
}
