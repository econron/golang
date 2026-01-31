package basestacks

import (
	"fmt"
	"time"
	"sync"
)

func updateFile3(i int) {
	fmt.Printf("Updating file %d\n", i)
	time.Sleep(1 * time.Second) // IO待ちの擬似的な再現
}

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1) // waitGroupにgoroutineの数を追加する
// 		go func(i int) {
// 			defer wg.Done() // goroutineの完了を通知する
// 			updateFile3(i)
// 		}(i)
// 	}
// 	wg.Done() // 全てのgoroutineの完了を待つ
// }


// waitGroupのポインタ型を引数として渡し、メソッド内部でwg.Doneする形式
func updateFile3v2(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Printf("Updating file %d\n", i)
	time.Sleep(1 * time.Second) // IO待ちの擬似的な再現
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1) // waitGroupにgoroutineの数を追加する
		go updateFile3v2(&wg, i)
	}
	wg.Done() // 全てのgoroutineの完了を待つ
}