package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	const (
		maxTurns = 5
		text     = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?`
	)
	source:=rand.NewSource(time.Now().UnixNano())
	rng:=rand.New(source)

	arg:=os.Args[1:]

	if len(arg)!=1{
		fmt.Println("Enter a valid number")
		return
	}
	guess,err:=strconv.Atoi(arg[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}
	if guess<0{
		fmt.Println("Enter a positive number")
		return
	}
	for i:=0;i<maxTurns;i++{
		n:=rng.Intn(guess)+1
		fmt.Println(n)
		if n==guess{
			fmt.Println("🎉  YOU WIN!")
			return
		}

	}
	fmt.Println("☠️  YOU LOST... Try again?")

}