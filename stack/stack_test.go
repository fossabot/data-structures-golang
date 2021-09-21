package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	var testCases = []struct {
		items []int
		want  Stack
	}{
		{
			[]int{100, 200, 300},
			Stack{[]int{100, 200, 300}},
		},
		{
			[]int{10, 20, 30},
			Stack{[]int{10, 20, 30}},
		},
	}

	for _, tc := range testCases {
		myStack := NewStack()
		for _, item := range tc.items {
			myStack.Push(item)
		}
		if !reflect.DeepEqual(myStack, tc.want) {
			t.Error("expected:", tc.want, "got:", myStack)
		}
	}
}

func TestPop(t *testing.T) {
	var testCases = []struct {
		stack Stack
		want  Stack
	}{
		{
			Stack{[]int{100, 200, 300}},
			Stack{[]int{100, 200}},
		},
		{
			Stack{[]int{10, 20, 30}},
			Stack{[]int{10, 20}},
		},
	}
	for _, tc := range testCases {
		tc.stack.Pop()
		if !reflect.DeepEqual(tc.stack, tc.want) {
			t.Error("expected:", tc.want, "got:", tc.stack)
		}
	}
}
