package main

import "fmt"

// TreeNode 定義
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFSによるレベル順走査
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root} // BFS 用のキューを初期化
	fmt.Printf("first queue: %#v\n", queue)

	for len(queue) > 0 {
		size := len(queue) // 現在のレベルにあるノードの数
		var level []int    // そのレベルの値を格納するスライス

		for i := 0; i < size; i++ {
			node := queue[0] // キューの先頭を取得
			fmt.Printf("node: %#v\n", node)
			queue = queue[1:] // キューから削除
			fmt.Printf("queue: %#v\n", queue)
			level = append(level, node.Val) // ノードの値を追加
			fmt.Printf("level: %#v\n", level)

			if node.Left != nil { // 左の子ノードをキューに追加
				queue = append(queue, node.Left)
			}
			if node.Right != nil { // 右の子ノードをキューに追加
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level) // レベルごとの結果を保存
	}

	return res
}

// テスト
func main() {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	result := levelOrder(root)
	fmt.Println(result) // [[3], [9, 20], [15, 7]]
}
