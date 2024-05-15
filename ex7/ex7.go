package main

import (
	"fmt"
	"unicode"
)

// count the number of letters and numbers in a string

func main(){
	ans:= count("Hello world! 123")
	fmt.Println(ans)
}
 
func count(str string) string{
	letter:=0
	num:=0
	for _,val:=range str{
		if(unicode.IsLetter(val)){
			letter++
		} else if(unicode.IsNumber(val)){
			num++
		} else{
			continue
		}
	}
	return fmt.Sprintf("Letters: %v & Numbers: %v",letter,num)
}