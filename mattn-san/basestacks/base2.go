package basestacks

import (
	"fmt"
	"time"
)

func updateFile2(i int) {
	fmt.Printf("Updating file %d\n", i)
	time.Sleep(1 * time.Second) // IO待ちの擬似的な再現
}

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go updateFile2(i)
// 	}
// }