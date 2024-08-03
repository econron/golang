package main

import (
	"fmt"
)

func main(){
	var a *string = nil
	// bに「型は*stringでnil値」という状態を渡しているので、bはnilと見做されない
	var b interface{} = a
	fmt.Println("a == nil:", a == nil)
	fmt.Println("b == nil:", b == nil)
	// a == b は値の比較
	fmt.Println("a == b", a == b)
}