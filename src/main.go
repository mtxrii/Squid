package main

import (
	"fmt"
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

				//if _, err := http.Get(url); err != nil {
				//	print("Error: this website cannot be found.")
				//	return
				//}

				//  fastest - request with the lowest ms
				//  slowest - request with the highest ms
				//  total - all ms added up, used for mean & mdn
				//  smallest & largest - their byte sizes
				//  successRate - number of times a result is found
				fastest, slowest, total, smallest, largest, successRate := 0, 0, 0, 0, 0, 0
				var mdn []int

				println("Sending " + strconv.Itoa(n) + " requests...")
				for i := 1; i <= n; i++ {
					print("Attempt " + strconv.Itoa(i) + " - ")
					speed, size, err := Ping(url)
					if err != nil {
						println("ERROR:   " + strconv.Itoa(speed))
						continue
					}

					println("SUCCESS: " + strconv.Itoa(speed) + " ms (" + strconv.Itoa(size) + " bytes)")
					if speed < fastest {
						fastest = speed
					}
					if speed > slowest {
						slowest = speed
					}
					if size < smallest {
						smallest = size
					}
					if size > largest {
						largest = size
					}
					mdn = append(mdn, speed)
					successRate++
					total += speed
				}

				println("\nSUMMARY:\n" +
					"Total requests -    " + strconv.Itoa(n) + " ms\n" +
					"Fastest response -  " + strconv.Itoa(fastest) + " ms\n" +
					"Slowest response -  " + strconv.Itoa(slowest) + " ms\n" +
					"Average response -  " + fmt.Sprintf("%.2f", float32(total)/float32(n)) + " ms\n" +
					"Median response -   " + strconv.Itoa(Median(mdn)) + " ms\n" +
					"Smallest response - " + strconv.Itoa(smallest) + " bytes\n" +
					"Largest response -  " + strconv.Itoa(largest) + " bytes")

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
