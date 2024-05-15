package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Write a program that accepts sequence of lines as input and prints the lines after making all characters in the sentence capitalized.

func main(){
	fmt.Print("Enter the string: ")
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		text:=scanner.Text()
		if(text=="q"){
			break;
		}
		fmt.Println(strings.ToUpper(text))
	}
}