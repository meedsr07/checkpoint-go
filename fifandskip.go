package main

import (
	"fmt"
	
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
func FifthAndSkip(str string) string {
if str == ""  {
	  return "\n"
}
if len(str) < 5 {
	return "invalid input\n"
}
res  := ""
for _, v := range str {
	if v != ' ' {
		res+= string(v)
	}
}
reslte := ""
count := 0
for i := 0; i <len(res); i++ {
	reslte += string(res[i])
	count++
	if count == 5{
		reslte += " "
		count = 0
		i++
	}
}
return  reslte + "\n"
}