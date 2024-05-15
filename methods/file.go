package main
import (
	"fmt"
)
type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user active:",u.Status)
}
func main(){
	mystruct:=User{"Harry","Harry@gmail.com",true,22}
	mystruct.GetStatus()
}