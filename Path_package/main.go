package main

import(
	"fmt"
	"path"
)

func main(){
	var dir,file string
	dir,file=path.Split("css/main.css")
	fmt.Println(dir)
	fmt.Println(file)
}