package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result []int
	if root != nil {
		result = append(result, root.Val)
	}

	rightSideViewAux(root, &result, 1)
	return result
}

func rightSideViewAux(current *TreeNode, result *[]int, level int) {
	if current == nil {
		return
	}

	if len(*result) < level {
		*result = append(*result, current.Val)
	}

	rightSideViewAux(current.Right, result, level+1)
	rightSideViewAux(current.Left, result, level+1)
}

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	}
	fmt.Println(rightSideView(tree1))

	tree2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(rightSideView(tree2))
}
