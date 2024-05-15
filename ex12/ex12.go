package main
import (
	"fmt"
)
type Item struct{
	name string
	id int
}
type Game struct{
	Item
	genre string
}
func main(){
	game:=[]Game{{
		Item:Item{"God of War",1234},
		genre:"action",
	},{
		Item:Item{"Spiderman",3553},
		genre: "adventure",
	}}

	for _,val:=range game{
		fmt.Println(val.Item.id)
		fmt.Println(val.genre)
	}
}