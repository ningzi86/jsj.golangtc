package arithmetic

import (
	"testing"
	"fmt"
)

type queue struct {
	arr []int
}

func (q *queue) pull() int {
	if len(q.arr) == 0 {
		panic("empty queue")
	}
	v := q.arr[0]
	q.arr = q.arr[1:]
	return v
}

func (q *queue) add(v int) {
	q.arr = append(q.arr, v)
}

func TestQueue(t *testing.T) {

	q := queue{}
	q.add(1)
	q.add(2)
	q.add(3)

	fmt.Println(q.pull())
	fmt.Println(q.pull())
	q.add(4)
	fmt.Println(q.pull())
	fmt.Println(q.pull())
	fmt.Println(q.pull())

}
