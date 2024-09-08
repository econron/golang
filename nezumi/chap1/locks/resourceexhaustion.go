package locks

import (
	"fmt"
	"sync"
	"time"
)

func ResourceExhaustion() {
	fmt.Println("this is resource exhaustion process.")

	var wg sync.WaitGroup
	var sharedlock sync.Mutex
	const runtime = 1*time.Second

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now();time.Since(begin) <= runtime; {
			sharedlock.Lock()
			time.Sleep(3*time.Nanosecond)
			sharedlock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now();time.Since(begin) <= runtime; {

			sharedlock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedlock.Unlock()

			sharedlock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedlock.Unlock()

			sharedlock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedlock.Unlock()

			count++
		}

		fmt.Printf("Police worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()
}