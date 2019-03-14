package arithmetic

import (
	"testing"
	"fmt"
)

type stack struct {
	arr []int
}

func (q *stack) pop() int {
	if len(q.arr) == 0 {
		panic("empty queue")
	}
	v := q.arr[len(q.arr)-1]
	q.arr = q.arr[:len(q.arr)-1]
	return v
}

func (q *stack) push(v int) {
	q.arr = append(q.arr, v)
}

func TestStack(t *testing.T) {

	st := &stack{}

	st.push(1)
	st.push(2)
	st.push(3)

	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st.pop())

}
