package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
			Usage:   "Makes an HTTP request to the specified URL and prints the response directly to the console.",
			Action: func(c *cli.Context) {
				url := c.Args().Get(0)
				if url == "" {
					print("Usage: url <full url>")
					return
				}

				unknown := "Error: this website cannot be found."
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
		{
			Name:    "profile",
			Aliases: []string{"ping"},
			Usage:   "Sends specified number of requests to the specified URL, times them, and delivers speed statistics.",
			Action: func(c *cli.Context) {
				times := c.Args().Get(0)
				url := c.Args().Get(1)
				if times == "" || url == "" {
					print("Usage: profile <number of requests> <full url>")
					return
				}

				n, err := strconv.Atoi(times)
				if err != nil {
					print("Error: number of requests must be a positive integer.")
					return
				}

				if n < 1 {
					print("Error: number of requests must be a positive integer.")
					return
				}

				if _, err := http.Get(url); err != nil {
					print("Error: this website cannot be found.")
					return
				}

				speed, err := Ping(url)
				if err != nil {
					print(err)
					return
				}
				print(speed)

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
