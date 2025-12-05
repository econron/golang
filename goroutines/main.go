package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// workerPoolAntiPattern()
	// workerPoolDeadlockPattern1()
	workerPoolPattern1()
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
		wg.Wait() // シェフ全員の仕事の終わりを待つ
		close(orders) // その後注文を閉じる
		fmt.Println("すべての料理が配膳されました！")
	}()

	<-done // ウェイターの仕事の終わりを待つ

	fmt.Println("すべての料理が配膳されました！")
}