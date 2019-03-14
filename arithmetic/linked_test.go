package arithmetic

import (
	"testing"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func (n *Node) append(node *Node) *Node {
	current := n
	for current.next != nil {
		current = current.next
	}
	current.next = node
	return n
}

func (n *Node) show() {

	current := n
	for {
		fmt.Println(current.data)
		if current.next == nil {
			break
		}
		current = current.next
	}

}

func (n *Node) delete() {
	if n.next == nil {
		return
	}
	n.next = n.next.next
}

func (n *Node) after(node *Node) {
	n.next, node.next = node, n.next
}

func TestNode(t *testing.T) {

	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}
	n4 := &Node{data: 4}
	n5 := &Node{data: 5}

	n1.append(n2).append(n3).append(n4).append(n5)
	n1.show()

	n1.delete()
	n1.show()

	n1.after(&Node{data: 100})
	n1.show()

}
