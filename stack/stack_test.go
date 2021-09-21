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
		name   string
		stack  Stack
		want   Stack
		popVal int
	}{
		{
			"Remove the value 300 from the stack",
			Stack{[]int{100, 200, 300}},
			Stack{[]int{100, 200}},
			300,
		},
		{
			"Remove the value 30 from the stack",
			Stack{[]int{10, 20, 30}},
			Stack{[]int{10, 20}},
			30,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pop := tc.stack.Pop()
			if !reflect.DeepEqual(tc.stack, tc.want) {
				t.Error("expected:", tc.want, "got:", tc.stack)
			}
			if pop != tc.popVal {
				t.Error("expected:", tc.popVal, "got:", pop)
			}
		})
	}
}
