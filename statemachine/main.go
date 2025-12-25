package main

import (
	"fmt"
	sm "statemachine/statemachine"
	ddd "statemachine/dddlike"
)

func main() {
	// statemachine愚直？パターン
	state := sm.NewStockState(sm.OutOfStock)
	newState, err := state.Transition(sm.StockArrived)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("new state is %#v\n", newState)

	// dddパターン
	book := ddd.NewBook("Valkey secrets")
	err = book.Arrived()
	if err != nil {
		fmt.Printf("error occured while transfering status. %#v", err)
	}
	fmt.Printf("book object: %#v\n", book)
	err = book.Sold()
	if err != nil {
		fmt.Printf("error occured while transfering status. %#v", err)
	}
	fmt.Printf("book object: %#v\n", book)
	// あえてエラーを起こす
	err = book.Sold()
	if err != nil {
		fmt.Printf("error occured while transfering status. %#v", err)
	}
}