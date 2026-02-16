package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
func WordFlip(str string) string {
if str == "" {
	return "invlid input\n"
}
var words []string
var word string
for i := 0; i < len(str); i++ {
	c:= (str[i])
	if c != ' ' {
		word += string(c)
	}else if word != "" {
		words = append(words, word)
		word = ""
	}
}
if word != "" {
	words = append(words,word)
}
res := ""
for i := len(words) -1; i>=0 ; i-- {
	res += string(words[i])
	if i != 0 {
		res += " "
	}
}
return  res + "\n"
}
