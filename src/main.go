package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "Squid"
	app.Usage = "A web getter & speed tester CLI"
	app.Author = "mtxrii"
	app.Version = "0.1.1"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "url",
			Aliases: []string{"fetch", "get"},
			Usage:   "Makes an HTTP request to the URL and prints the response directly to the console.",
			Action: func(c *cli.Context) {
				url := c.Args().Get(0)
				if url == "" {
					print("Usage: url <full url>")
					return
				}

				unknown := "Sorry, but this website cannot be found"
				resp, err := http.Get(url)
				if err != nil {
					print(unknown)
					return
				}

				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					print(unknown)
					return
				}

				print("\n" + string(body))
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
