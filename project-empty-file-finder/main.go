package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg)==0{
		fmt.Println("Provide a directory")
		return 
	}

	files,err:=os.ReadDir(arg[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _,file:= range files{
		info,err:=file.Info()
		if err!=nil{
			fmt.Println(err)
			return
		}
		if info.Size()>0{
			name:=file.Name()
			fmt.Println(name)
		}
	}
}