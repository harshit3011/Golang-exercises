package main

import (
	"fmt"
	// "io"
	"log"
	// "net/http"
	"os"
)

func main() {
	// file,err:=os.Create("novel.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// resp,err:=http.Get("https://www.gutenberg.org/cache/epub/73881/pg73881.txt")
	// if err!=nil{
	// log.Fatal(err)
	// }
	// defer resp.Body.Close()
	// n,err:=io.Copy(file,resp.Body)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// log.Printf("Successfully copied %v bytes from the URL ",n)

	file,err:=os.Open("./novel.txt")
	if err != nil {
		log.Fatal(err)
	}
	data:=make([]byte,10000)
	count,err:=file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %v bytes:\n %q\n",count,data[:count])
}