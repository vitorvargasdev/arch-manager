package cli

import (
	cli "github.com/urfave/cli/v2"
	installArch "github.com/vitorvargasdev/archmanager/pkgs/install"
)

func Prepare() *cli.App {
	app := cli.NewApp()
	app.Name = "Arch Manager ~ Developed by @vitorvargasdev"
	app.Usage = "Installation and management of Arch Linux with eas"

	Flags := []cli.Flag{}

	app.Commands = []*cli.Command{
		{
			Name:   "install",
			Usage:  "Install Arch Linux with Arch Manager",
			Action: install,
			Flags:  Flags,
		},
	}

	return app
}

func install(c *cli.Context) error {
	installArch.Run()

	return nil
}
