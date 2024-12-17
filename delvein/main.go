package main

import "fmt"

func main() {
    x := 10
    y := 20
    result := add(x, y)
    fmt.Println("Result:", result)
}

func add(a, b int) int {
    return a + b
}
