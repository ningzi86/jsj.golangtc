package tests

import (
	"testing"

	"fmt"

	. "github.com/ahmetalpbalkan/go-linq"
)

func Test_Aggregate(t *testing.T) {

	tests := []struct {
		input interface{}
		want  interface{}
	}{
		{[]string{"abc", "ab", "a", "e", "grape"}, "passionfruit"},
		{[]string{}, nil},
	}

	result := From(tests[0].input).Aggregate(func(r interface{}, i interface{}) interface{} {

		if len(r.(string)) > len(i.(string)) {
			return i
		}
		return r

	})

	fmt.Println(result.(string))

}

func TestAggregateWithSeed(t *testing.T) {
	input := []string{"apple", "mango", "orange", "banana", "grape"}
	want := "passionfruit"

	r := From(input).AggregateWithSeed(want,
		func(r interface{}, i interface{}) interface{} {
			fmt.Println(r, i)
			if len(r.(string)) > len(i.(string)) {
				return r
			}
			return i
		})

	fmt.Println(r)

	if r != want {
		t.Errorf("From(%v).AggregateWithSeed()=%v expected %v", input, r, want)
	}

	
}



