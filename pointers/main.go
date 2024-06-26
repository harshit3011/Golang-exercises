package main

import "fmt"

type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func (u User) UpdateEmail(email string){
	u.email=email
}
func (u *User) UpdatePoint(email string){
	u.email=email
}

func main() {
	u := User{
		email:    "cool@gmail.com",
		username: "Harry",
		age:      22,
	}
	u.UpdatePoint("wow@gmail.com")
	fmt.Println(u.Email())
}