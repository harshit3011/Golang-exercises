package main
import (
	"fmt"
)
//sum of n numbers
func main(){
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&n)
	sum:=(n*(n+1))/2
	fmt.Println("Sum is", sum)
}