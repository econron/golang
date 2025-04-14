package main

import (
	"fmt"
)

// backtrack は、再帰的に順列を構築して result に追加します。
func backtrack(nums []int, path []int, used []bool, result *[][]int) {
	// 順列が完成したら、path のコピーを result に追加する
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// 各要素について、使用していないなら順列に追加
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		
		fmt.Println("before backtrack")
		fmt.Printf("index is %v", i)
		fmt.Println(path)
		backtrack(nums, path, used, result)
		// バックトラック: 直前の選択を取り消す
		fmt.Println("after backtrack")
		path = path[:len(path)-1]
		fmt.Printf("index is %v", i)
		fmt.Println(path)
		used[i] = false
	}
}

// permute は、入力された nums のすべての順列を返します。
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, used, &result)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	permutations := permute(nums)
	fmt.Println(permutations)
}
