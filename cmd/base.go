package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

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

	c = append(c, cli.Command{
		Name:     "docker-clean",
		Category: "Docker",
		Usage:    "Cleans docker enviroment",
		Action: func(c *cli.Context) {
			cmdContainers := exec.Command("docker", "ps", "-q", "--no-trunc")
			var out bytes.Buffer
			cmdContainers.Stdout = &out
			errContainer := cmdContainers.Run()
			if errContainer != nil {
				fmt.Println("No docker images found")
				return
			}
			commandStop := exec.Command("docker", "stop")
			commandStop.Run()
			commandRmv := exec.Command("docker", "rm", out.String())
			outputRmv, erroRmv := commandRmv.CombinedOutput()
			if erroRmv != nil {
				fmt.Println(fmt.Sprint(erroRmv) + ": " + string(outputRmv))
				return
			}
			commandPrune := exec.Command("docker", "network", "prune")
			outputPrune, erroPrune := commandPrune.CombinedOutput()
			if erroPrune != nil {
				fmt.Println(fmt.Sprint(erroPrune) + ": " + string(outputPrune))
				return
			}
			fmt.Println("Docker is now clean")

		},
	})

	return c
}
