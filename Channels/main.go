package main

import (
	"fmt"

	"github.com/harshit3011/myniceprogram/helpers"
)

const numPool= 100

func Calculate(intChan chan int){
	random:=helpers.RandomNumber(numPool)
	intChan <- random

}

func main(){
	intChan:= make(chan int)
	defer close(intChan)
	go Calculate(intChan)

	num:= <- intChan
	fmt.Println(num)
}