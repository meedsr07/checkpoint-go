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
	for i := 0; i <len(str); i++ {
		c := str[i]
		if c != ' '{
			word += string(c)
		}else if word != "" {
			words = append(words,word)
			word = ""
		}
	} 
	if word != "" {
		words = append(words, word)
	}
	for i := len(words)-1; i >=0; i-- {
		Printing(words[i])
		if i != 0{
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
func Printing(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}