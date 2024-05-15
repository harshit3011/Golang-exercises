package main

import(
	"fmt"
)

//With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value


func main(){
	var n int
	fmt.Print("Enter the number - ")
	fmt.Scanln(&n)
	ans:=result(n)
	fmt.Println("The required map is: ")
	fmt.Println(ans)
}

func result(n int) map[int]int{
	res:=make(map[int]int)
	for i:=1;i<=n;i++{
		res[i]=i*i
	}
	return res
}