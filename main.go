package main

import (
	"fmt"
	"time"
)

func main() {
	typeThis("Hello World")
}

func typeThis(text string) {
	for _, c := range text {
		fmt.Print(string(c))
		if string(c) == string(' ') {
			time.Sleep(30 * time.Millisecond)
		} else {
			time.Sleep(85 * time.Millisecond)
		}
	}
}
