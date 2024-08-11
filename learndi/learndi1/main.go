package main

import (
	"fmt"
	"learndi/usecase"
	"learndi/repository"
)

func main() {
	r := repository.New()
	u := usecase.New(r)
	fmt.Println(u.GetUser())
}