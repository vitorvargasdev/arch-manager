package main

import (
	"log"
	"os"

	"github.com/vitorvargasdev/archmanager/pkgs/cli"
)

func main() {
	cli := cli.Prepare()

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
