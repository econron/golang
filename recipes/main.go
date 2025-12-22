package main 

import (
	"iter"
	"fmt"
	"recipes/stringrecipes"
)

// 構文
// func YourIterator() iter.Seq2[K,V] {
// 	return func(yield func(K,V) bool) {
// 		for i,v := range someData {
// 			if !yield(i,v) {
// 				return
// 			}
// 		}
// 	}
// }

func Backward[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

// 問題1：偶数とそのインデックスを返す
// スライスを受け取りその中の偶数だけを元のインデックスと一緒に返すイテレータ

// iter.Seq2[int,int]の場合はyieldは2つのintを引数に取る必要がある
// 今回返したいのは[int,int]なのでそれを戻り値に設定した
func OnlyEvens(nums []int) iter.Seq2[int,int] {
	return func(yield func(int, int) bool) {
		for i := 0; i < len(nums); i++ {
			if nums[i]%2 == 0 {
				// 配列から値を取得する本体の処理だとみなす
				if !yield(i, nums[i]) {
					return
				}
			}
		}
	}
}

// 問題2：スライスのチャンク化
// チャンク化をイテレータで実装する
func Chunks[E any](s []E, size int) iter.Seq2[int, []E] {
	return func(yield func(int, []E) bool) {
		chunkIdx := 0
		// チャンクサイズずつ増やしていけば良い
		for i := 0; i < len(s); i+= size {
			end := i + size
			if end > len(s) {
				end = len(s)
			}
			// ここがメイン処理
			// 新しいスライスを作る前に元のスライスを切り出せないか考えるべき
			chunk := s[i:end]
			if !yield(chunkIdx, chunk) {
				return
			}
			chunkIdx++
		}
	}
}

// 問３　mapのフィルタリング
func FilterMap(m map[string]int, threshold int) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for k,v := range m {
			if v < threshold {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func main() {
	nums := []string{"Go", "Rust", "Zig"}
	for i, val := range Backward(nums) {
		fmt.Printf("%d: %s\n", i, val)
	}
	nums2 := []int{10,11,12,13,14}
	for i, v := range OnlyEvens(nums2) {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	stringrecipes.Basic1()
	stringrecipes.Basic2()
	stringrecipes.Basic3()
	stringrecipes.Basic4()
	fmt.Printf("RLE: %v\n", stringrecipes.RLE("a3b2c4"))
}