package practice

import (
	"fmt"
	// "sync"
)

func Prob2() {
	ch := make(chan string, 2)
	go func() {
		ch <- "Hello World"
	}()
	fmt.Printf("%s", <-ch)
}