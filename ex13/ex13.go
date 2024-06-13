//basic web app

package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter,r* http.Request){
	fmt.Fprintf(w, "Whoa, Go is Neat!!")
}

func about_handler(w http.ResponseWriter,r* http.Request){
	fmt.Fprintf(w,`<h1>Hey There</h1>
		<p>Go is fast!</p>
		 <p>....and simple!</p>
		 <p>You %s even add %s</p>`,"can","<strong>variables</strong>")
}

func main(){
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about",about_handler)
	fmt.Println("Server is starting")
	http.ListenAndServe(":8000",nil)
}