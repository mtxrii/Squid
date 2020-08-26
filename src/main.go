package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	fmt.Print(encode("hello world"))
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

func encode(text string) string {
	var charGlyphs = GetCharGlyphs()
	var numberGlyphs = GetNumberGlyphs()
	var specialGlyphs = GetSpecialGlyphs()

	var result string
	for _, c := range text {
		if c == ' ' {
			result += string(' ')
		} else if unicode.IsDigit(c) {
			result += numberGlyphs[c]
		} else if unicode.IsUpper(c) {
			result += "$" + charGlyphs[unicode.ToLower(c)]
		} else if unicode.IsLower(c) {
			result += charGlyphs[c]
		} else {
			key, exists := specialGlyphs[c]
			if exists {
				result += key
			} else {
				result += string(c)
			}
		}
	}
	return result
}
