package main

import (
	"fmt"
	"math/rand"
	"time"
)

//

func generatePassword(length int, charset string) string{
	rand.Seed(time.Now().UnixNano())
	password:=""
	for i:=0;i<length;i++{
		password+=string(charset[rand.Intn(len(charset))])
	}
	return password
}

func main(){
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	var len int
	fmt.Print("Enter the length of password: ")
	fmt.Scanln(&len)

	ans:=generatePassword(len,charset)

	fmt.Printf("The generated password is: %v",ans)

}