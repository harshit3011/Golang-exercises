package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Write a program, which will find all such numbers between 100 and 300 (both included) such that each digit of the number is an even number. The numbers obtained should be printed in a comma-separated sequence on a single line.
func main(){

	var start int64=100
	var end int64=300
	ans:=[]string{}
	for i:=start;i<=end;i++{
		num:=strconv.FormatInt(i,10)
		a:=int(num[0])
		b:=int(num[1])
		c:=int(num[2])

		if a%2==0 && b%2==0 && c%2==0{
			ans = append(ans, num)
		}
	}
	result:= strings.Join(ans,",")
	fmt.Println(result)
}