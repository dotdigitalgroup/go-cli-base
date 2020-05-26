package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func SetupBase(c []cli.Command) []cli.Command {
	c = append(c, cli.Command{
		Name:    "base",
		Aliases: []string{"b"},
		Usage:   "Show a hello world command line",
		Action: func(c *cli.Context) {
			fmt.Println("Hello World")
		},
	})

	return c
}
