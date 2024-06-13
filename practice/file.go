package main
import (
	"fmt"
)
// func equals(x,y map[int]string) bool{
// 	if(len(x)!=len(y)){
// 		return false
// 	}
// 	for k,xv:=range x{
// 		yv,ok:=y[k]
// 		if(!ok || yv!=xv){
// 			return false
// 		}
// 	}
// 	return true
// }
// func main(){
// 	x:=make(map[int]string)
// 	y:=make(map[int]string)
// 	x[0]="Cool"
// 	x[1]="Name"
// 	x[2]="Wowee"
// 	y[0]="Cool"
// 	y[1]="Name"
// 	y[2]="Wowee"
	
// 	fmt.Println("Are the two maps equal:",equals(x,y))
// }


type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct{
	Name string
	Breed string
}

type Lion struct{
	Name string
	Place string
}

func (d Dog) Says() string{
	return "woof woof"
}
func (l Lion) Says() string{
	return "Rrarrrr"
}

func (d Dog) NumberOfLegs() int{
	return 4
}
func (l Lion) NumberOfLegs() int{
	return 4
}

func main(){
	dog:= Dog{Name: "Spike",Breed: "Bulldog"}
	lion:=Lion{Name: "Simba", Place: "Kenya"}

	fmt.Println(dog.Says())
	fmt.Println(dog.NumberOfLegs())
	fmt.Println(lion.Says())
	fmt.Println(lion.NumberOfLegs())
}