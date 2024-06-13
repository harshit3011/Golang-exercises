package main

import (
	"fmt"

	"github.com/harshit3011/myniceprogram/helpers"
)
func main(){
	var myType helpers.SomeType
	myType.TypeName="Harry"
	myType.TypeNum=22
	fmt.Print(myType.TypeName," is ",myType.TypeNum," years old")
}