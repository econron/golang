package practice

import (
	"fmt"
	"sync"
)

func Prob1() {
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("print! %d\n", i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("DONE")
}