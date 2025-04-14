package main

import (
	"fmt"
	"log"
)

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 // 各要素は最小1つの部分列を形成する
	}

	maxLIS := 1

	for i := 1; i < n; i++ {
		log.Printf("i: %d\n", i)
		for j := 0; j < i; j++ {
			log.Printf("j: %d\n", j)
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				log.Printf("dp[%d]: %d\n", i, dp[i])
			}
		}
		maxLIS = max(maxLIS, dp[i])
	}

	return maxLIS
}

// 最大値を求める関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums)) // 4
}
