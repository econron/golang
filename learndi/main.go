package main

import (
	"fmt"
	"learndi/usecase"
)

func main() {
	u := usecase.New()
	fmt.Println(u.GetUser())
}