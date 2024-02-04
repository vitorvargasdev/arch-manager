package cli

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func Prepare() *cli.App {
	app := cli.NewApp()
	app.Name = "Arch Manager ~ Developed by @vitorvargasdev"
	app.Usage = "Installation and management of Arch Linux with easy"

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
	fmt.Println("Installing Arch Linux with Arch Manager")

	return nil
}
