package main

import "fmt"

func main() {
	nodes := newListNodes([]int{1,2,3,4,5,6})
	if isVisited(nodes) {
		fmt.Println("visited!")
	}
	fmt.Println("non visited!")
}

type ListNode struct {
	Val int
	Next *ListNode
}

func newListNodes(values []int) *ListNode {
	pointer := &ListNode{
		Val: values[0],
		Next: nil,
	}
	for i, v := range values {
		if i == 0 {
			continue
		}
		node := &ListNode{
			Val: v,
			Next: pointer,
		}
		pointer = node
	}
	return pointer
}

func isVisited(n *ListNode) bool {
	visited_nodes := make(map[*ListNode]bool)
	current_node := n
	for current_node != nil {
		if visited_nodes[current_node] {
			return true
		}
		visited_nodes[current_node] = true
		current_node = current_node.Next
	}

	fmt.Printf("%#v", visited_nodes)

	return false
}