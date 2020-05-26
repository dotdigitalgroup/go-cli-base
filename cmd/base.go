package cmd

import (
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
			command := exec.Command("docker", "stop $(docker ps -aq) && docker rm $(docker ps -qa) && docker network prune -f")
			output, erro := command.CombinedOutput()
			if erro != nil {
				fmt.Println(fmt.Sprint(erro) + ": " + string(output))
				return
			}
			fmt.Println("Docker is now clean")

		},
	})

	return c
}
