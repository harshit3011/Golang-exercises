package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuff": {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}
	arg := os.Args[1:]
	if len(arg)!=1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}
	delete(houses,"bobo")
	key:=arg[0]
	if value,ok:=houses[key];ok{
		sort.Strings(value)
		fmt.Printf("~~~ %s students ~~~",key)
		fmt.Println()
		fmt.Println()
		for _,v:=range value{
			fmt.Printf("+ %s\n",v)
		}
	} else {
		fmt.Printf("Sorry. I don't know anything about %s.",key)
		return
	}
}