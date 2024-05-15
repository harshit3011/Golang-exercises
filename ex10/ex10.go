package main

import (
	"fmt"
	"math/rand"
	"time"
)

//guess the number
func main(){
	rand.Seed(time.Now().UnixNano())
	var guess int
	fmt.Print("Enter the number to be guessed: ")
	fmt.Scanln(&guess)
	
	tries:=10
	flag:=false
	for i := 0; i < tries; i++ {
		n:=rand.Intn(guess+1)
		fmt.Print(n," ")
		if(n==guess){
			flag=true
			fmt.Println("Correct number!!")
			break
		}
	}
	if !flag{
		fmt.Println("Correct Number could not be guessed")
	}
}