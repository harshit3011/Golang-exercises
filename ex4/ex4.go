package main

import (
	"fmt"
	"strconv"
	"strings"
)

//take a string of comma-separated numbers and return a slice of int

func main(){
	var str string
	fmt.Print("Enter the string: ")
	fmt.Scanln(&str)
	ans:= result(str)
	fmt.Println(ans)
}
func result(s string) []int{
	intgs:=strings.Split(s, ",")
	res:=make([]int,len(intgs))
	for i,val:=range intgs{
		res[i],_=strconv.Atoi(val)
	}
	return res
}