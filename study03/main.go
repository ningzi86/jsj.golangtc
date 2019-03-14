package main

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (n TreeNode) Print() {
	fmt.Println(n.Value)
}

func (n *TreeNode) TraverseFunc(f func(n *TreeNode)) {

	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}

func main() {

	node := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 3,
			},
		},
		Right: &TreeNode{
			Value: 4,
			Left: &TreeNode{
				Value: 5,
			},
		},
	}

	count := 0
	node.TraverseFunc(func(n *TreeNode) {
		count ++
		n.Print()
	})

	fmt.Println("count", count)

}
