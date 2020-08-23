package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Println(strconv.Itoa(i) + ")" + string(i))
	}

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
