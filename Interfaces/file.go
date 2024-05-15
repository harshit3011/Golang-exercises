package main
import (
	"fmt"
	"strings"
)
// type tank interface{
// 	area() float64
// 	vol() float64
// }
// type Dimensions struct{
// 	radius float64
// 	height float64
// }
// func (d Dimensions) area() float64{
// 	return (2*d.radius*d.height) + (2*3.14*float64(d.radius)*float64(d.height))
// }
// func (d Dimensions) vol() float64{
// 	return 3.14*d.radius*d.radius*d.height
// }
// func main(){
// 	var obj tank
// 	obj=Dimensions{12,20}
// 	fmt.Println("Area is: ",obj.area())
// 	fmt.Println("Vol is: ",obj.vol())
// }

type vowels interface{
	findVowels() []string
}
type Mystring string

func (str Mystring) findVowels() []string{
	ans:=[]string{}
	str=Mystring(strings.ToLower(string(str)))
	for _,val:= range str{
		if val=='a'||val=='e'||val=='i'||val=='o'||val=='u'{
			ans=append(ans,string(val))
		}
	}
	return ans
}
func main(){
	v:=Mystring("I am a good boy")
	fmt.Println(v.findVowels())
}