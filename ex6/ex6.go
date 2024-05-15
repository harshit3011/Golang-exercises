package main

import (
	"fmt"
	"sort"
	"strings"
)

// Accept a String separated by whitespace.
// input gets sorted and duplicates are removed

func main(){
	ans:= result("cool cool boy boy yay wow")
	fmt.Println(ans)
}

func result(str string) string{
	spl:=strings.Split(str, " ")
	m:=make(map[string]int)
	for _,val:=range spl{
		m[val]=m[val]+1
	}
	ans:=[]string{}
	for key,_:=range m{
		ans = append(ans, key)
	}
	sort.Strings(ans)
	ret:=strings.Join(ans," ")
	return ret
}