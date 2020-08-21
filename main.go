package main

import "fmt"
import "time"

func main() {
	typeThis("Hello World")
}


func typeThis(text string) {
	for _, c := range text {
		fmt.Print(string(c))
		time.Sleep(70 * time.Millisecond)
	}
}