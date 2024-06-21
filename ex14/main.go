package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l:=len(msg)

	s:= msg+ " "+strings.Repeat("!",l)
	fmt.Println(s)
}