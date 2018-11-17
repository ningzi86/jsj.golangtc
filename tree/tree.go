package main

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {

	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func (node *Node) Traverse2() {

	node.TraverseFunc(func(n *Node) {
		n.Print()
	})

	fmt.Println()

}

func (node *Node) TraverseFunc(f func(n *Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)

}

func CreateNode(value int) *Node {
	node := &Node{value, nil, nil}
	return node
}

func main() {

	root := Node{Value: 3}

	root.Left = &Node{}
	root.Left.Right = CreateNode(2)
	root.Left.Left = CreateNode(1)

	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Right.Left.SetValue(4)

	root.Traverse2()

	count := 0
	root.TraverseFunc(func(n *Node) {
		n.Value++
		count++
	})
	fmt.Println("total count", count)

	root.Traverse2()
}
