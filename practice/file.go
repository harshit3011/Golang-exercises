package main
import (
	"fmt"
)
func equals(x,y map[int]string) bool{
	if(len(x)!=len(y)){
		return false
	}
	for k,xv:=range x{
		yv,ok:=y[k]
		if(!ok || yv!=xv){
			return false
		}
	}
	return true
}
func main(){
	x:=make(map[int]string)
	y:=make(map[int]string)
	x[0]="Cool"
	x[1]="Name"
	x[2]="Wowee"
	y[0]="Cool"
	y[1]="Name"
	y[2]="Wowee"
	
	fmt.Println("Are the two maps equal:",equals(x,y))
}