package main

import (
	"fmt"
	"time"
)

func main() {
	encode("hello world")
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

func encode(text string) {
	var result string
	for _, c := range text {
		var char = string(c)
		if char == string(' ') {
			result += string(' ')
		} else {

		}
	}
}