package main

import (
	"go-cli-base/cmd"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {
	c := []cli.Command{}
	c = cmd.SetupBase(c)
	app.Commands = c
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
