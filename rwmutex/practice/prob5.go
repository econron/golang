package practice

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Prob5() {
	ch := make(chan string, 1)
	time.Sleep(time.Second * 3)
	// ctx := context.TODO()
	ch <- context.Cancel()

	go func() {
		for {
			fmt.Println("Running")
		}
	}()
}