package datas

import (
	"strings"
)

// 引数として渡された文字列を、スタックを使って逆順にして返してください。
func reverseString(s string) string {
	if len(s) == 0 {
		return s
	}
	stack := []rune{}
	for _, v := range s {
		stack = append(stack, v)
	}
	var builder strings.Builder
	for len(stack) != 0 {
		// 末尾から取得
		val := stack[len(stack)-1]
		// popしたので消す
		stack = stack[:len(stack)-1]
		// string builderで文字列を構築する
		builder.WriteRune(val)
	}
	return builder.String()
}

// スライスで与えられたタスク名（文字列）を、順番にキューに追加（Enqueue）し、その後、先頭から1つずつ取り出して（Dequeue）「完了済みスライス」に移動させてください。
func processTasks(tasks []string) []string {
	if len(tasks) == 0 {
		return tasks
	}
	queue := []string{}
	for _, v := range tasks {
		queue = append(queue, v)
	}
	completed := []string{}
	for len(queue) != 0 {
		elm := queue[0]
		queue[0] = ""
		queue = queue[1:]
		completed = append(completed, elm)
	}
	return completed
}