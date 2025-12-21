package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"goroutines/nezumi"
)

func main() {
	// workerPoolAntiPattern()
	// workerPoolDeadlockPattern1()
	// workerPoolPattern1()
	// scatterGatherPattern()
	endRandStream()
	nezumi.BadErrorHandling()
	errorResults := nezumi.GoodErrorHandling()
	fmt.Printf("Error Results: %v\n", errorResults.Errors)
	nezumi.PipelinePattern()
}

// worker pool pattern：複数の料理人が作ったら1人のウェイターができた順にどんどん処理していくイメージ
// ただ、これは同じwaitgroupを見ているせいで何もせずに処理が終了する。
func workerPoolAntiPattern() {
	// シェフが料理を渡し、ウェイターが料理を受け取るためのチャネル
	orders := make(chan string)

	var wg sync.WaitGroup

	// シェフが調理する
	chefs := []string{"chef1", "chef2", "chef3"}
	for _, chef := range chefs {
		wg.Add(1)
		go func(chefName string) {
			defer wg.Done()
			// 3皿ずつ作る
			for i := range 3 {
				dish := fmt.Sprintf("%s made dish %d", chefName, i+1)
				fmt.Printf("👨‍🍳 %s: %sを作成中...\n", chefName, dish)
				time.Sleep(1 * time.Second)
				orders <- dish
				fmt.Printf("🔔 %s: %sを配膳台に置いた\n", chefName, dish)
			}

		}(chef)
	}

	// ウェイターは料理を受け取ったそばから配膳していく
	// wg.Add(1)
	go func() {
		defer wg.Done()
		for dish := range orders {
			fmt.Printf("---------- 🤵 ウェイター: %sをお客さんに配膳しました\n", dish)
			fmt.Printf("---------- 🤵 お客さん: 料理を食べました\n")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		wg.Wait()
		close(orders)
	}()
	fmt.Println("すべての料理が配膳されました！")
}

// worker pool deadlock pattern
// 受信側のループに終了を伝えないと、チャネルが永遠に閉じず、デッドロックする
func workerPoolDeadlockPattern1() {
	// シェフが料理を渡し、ウェイターが料理を受け取るためのチャネル
	orders := make(chan string)

	var wg sync.WaitGroup

	// シェフが調理する
	chefs := []string{"chef1", "chef2", "chef3"}
	for _, chef := range chefs {
		wg.Add(1)
		go func(chefName string) {
			defer wg.Done()
			// 3皿ずつ作る
			for i := range 3 {
				dish := fmt.Sprintf("%s made dish %d", chefName, i+1)
				fmt.Printf("👨‍🍳 %s: %sを作成中...\n", chefName, dish)
				time.Sleep(1 * time.Second)
				orders <- dish
				fmt.Printf("🔔 %s: %sを配膳台に置いた\n", chefName, dish)
			}

		}(chef)
	}

	// ウェイターは料理を受け取ったそばから配膳していく
	for dish := range orders {
		fmt.Printf("---------- 🤵 ウェイター: %sをお客さんに配膳しました\n", dish)
		fmt.Printf("---------- 🤵 お客さん: 料理を食べました\n")
		time.Sleep(200 * time.Millisecond)
	}

	wg.Wait()
	close(orders)

	fmt.Println("すべての料理が配膳されました！")
}

// シェフ3人の調理をgoroutineにする
func workerPoolPattern1() {
	orders := make(chan string)

	var wg sync.WaitGroup

	// シェフが調理する
	chefs := []string{"chef1", "chef2", "chef3"}
	for _, chef := range chefs {
		wg.Add(1)
		go func(chefName string) {
			defer wg.Done()
			// 3皿ずつ作る
			for i := range 3 {
				dish := fmt.Sprintf("%s made dish %d", chefName, i+1)
				fmt.Printf("👨‍🍳 %s: %sを作成中...\n", chefName, dish)
				time.Sleep(1 * time.Second)
				orders <- dish
				fmt.Printf("🔔 %s: %sを配膳台に置いた\n", chefName, dish)
			}
		}(chef)
	}

	done := make(chan bool)
	go func() {
		for dish := range orders {
			fmt.Printf("---------- 🤵 ウェイター: %sをお客さんに配膳しました\n", dish)
			fmt.Printf("---------- 🤵 お客さん: 料理を食べました\n")
			time.Sleep(200 * time.Millisecond)
		}
		done <- true
	}()

	go func() {
		wg.Wait()     // シェフ全員の仕事の終わりを待つ
		close(orders) // その後注文を閉じる
		fmt.Println("すべての料理が配膳されました！")
	}()

	<-done // ウェイターの仕事の終わりを待つ

	fmt.Println("すべての料理が配膳されました！")
}

type Result struct {
	URL      string
	Duration time.Duration
	Data     string
}

func searchAPI(target string, duration time.Duration, ch chan<- Result) {
	fmt.Printf("Searching %s for %v\n", target, duration)
	time.Sleep(duration)
	ch <- Result{URL: target, Duration: duration, Data: fmt.Sprintf("data for %s", target)}
	fmt.Printf("Found %s in %v\n", target, duration)
}
func scatterGatherPattern() {
	start := time.Now()
	// バッファを3つ用意する
	results := make(chan Result, 3)
	go searchAPI("航空会社API", 1*time.Second, results)
	go searchAPI("天気API", 2*time.Second, results)
	go searchAPI("地図API", 3*time.Second, results)

	// バッファチャネルを利用しているのでいつ終わるかがわかっている
	// よってcloseしなくて良い
	for range 3 {
		result := <-results
		fmt.Printf("受信: %s\n", result.Data)
	}

	fmt.Printf("Time taken: %v\n", time.Since(start))
}

// TODO: 3つのAPI結果が揃うまで待ってからやる

// go言語による並行処理のコード
func endRandStream() {
	done := make(chan interface{})
	finished := make(chan bool)
	// doneチャネルを渡して、乱数生成を終了させる
	randStream := newRandStream(done, finished)
	for i := range 10 {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	// time.Sleep(1 * time.Second) 元のコードはこうだが、time.Sleepに依存するとgorotineの終了を確実に保証できない
	// よってfinishedチャネルを待つことでgoroutineの終了を確実に保証する
	<-finished
}

// 乱数を生成するgoroutine
func newRandStream(done <-chan interface{}, finished chan<- bool) <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStream done")
		defer close(randStream)
		for {
			select {
			case randStream <- rand.Int():
			case <-done:
				finished <- true
				return
			}
		}
	}()
	return randStream
}
