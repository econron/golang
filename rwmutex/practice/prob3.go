package practice

import (
	"fmt"
	"sync"
)

func Prob3() {
	size := 3
	ch := make(chan string, size)
	var wg sync.WaitGroup
	for i := range size + 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			str := fmt.Sprintf("string: %d\n", i)
			ch <- str
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	
	for str := range ch {
		fmt.Printf("strings: %s", str)
	}
}