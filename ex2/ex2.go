package main

import (
	"fmt"
)
 //Write a program which calculates the factorial of a given number

func main(){
	var num int
	fmt.Print("Enter a number")
	fmt.Scanln(&num)

	res:=fact(num)
	fmt.Println("Factorial of", num , "is", res)
}
func fact(num int) int{
	if num==0{
		return 0 
	}
	if num==1{
		return 1
	}
	return num*fact(num-1)
}





