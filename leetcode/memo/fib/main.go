package main

import "log"

func fib(n int) int {
	memo := [1000]int{0,1}
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func main() {
	log.Printf("%d", fib(10))
}