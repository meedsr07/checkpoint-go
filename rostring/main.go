package main

import (

	"os"


	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	var words []string
	word := ""
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c != ' ' {
			word += string(c)
		}else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	for i := 1; i < len(words); i++ {
		Print(words[i])
		z01.PrintRune(' ')
	}
	Print(words[0])
	z01.PrintRune('\n')
}
func Print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}