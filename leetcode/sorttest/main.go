package main

import (
	"fmt"
	"time"
	"math/rand"
	"sort"
)

func simpleArithmetic() {
	N := 1_000_000_000
	a := 0

	start := time.Now()
	for i := 0; i < N; i++ {
		a += i
	}
	elapsed := time.Since(start)

	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func sorts() {
	N := 10_000_000 // 10^7

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}

	start := time.Now()
	sort.Ints(arr)
	elapsed := time.Since(start)

	fmt.Printf("Sorting %d integers took %v\n", N, elapsed)
}

func main() {
	simpleArithmetic()
	sorts()
}