package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "Squid"
	app.Usage = "A web getter & speed tester CLI"
	app.Author = "mtxrii"
	app.Version = "0.1.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "url",
			Aliases: []string{"--fetch", "--get"},
			Usage:   "Makes an HTTP request to the URL and prints the response directly to the console.",
			Action: func(c *cli.Context) {
				print(c.Args().Get(0))
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
