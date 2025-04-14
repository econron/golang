package practice

import (
	"fmt"
	// "sync"
	"time"
)

func Prob4() {
	queue := make(chan string, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			msg := fmt.Sprintf("Task %d", i)
			fmt.Println("Enqueue:", msg)
			queue <- msg
			time.Sleep(time.Millisecond * 500)
		}
		close(queue) // これがないとrange queueが無限に待機する
	}()

	for msg := range queue {
		fmt.Println("Dequeue:", msg)
		time.Sleep(time.Second)
	}
	
	fmt.Println("Queue processing finished")
}
