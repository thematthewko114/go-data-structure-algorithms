package stacks

import (
	"fmt"
)

type Stack struct {
	values []string
	top    int
}

func NewStack() *Stack {
	return &Stack{
		values: make([]string, 0),
		top:    -1,
	}
}

func (st *Stack) Push(value string) {
	st.top++
	st.values = append(st.values, value)
}

func (st *Stack) Pop() string {
	if st.IsEmpty() {
		return ""
	}
	value := st.values[st.top]
	st.top--
	return value
}

func (st *Stack) IsEmpty() bool {
	return st.top == -1
}

func (st *Stack) PrintStack() {
	fmt.Println("Top")
	fmt.Println("------")
	for i := st.top; i >= 0; i-- {
		fmt.Println(st.values[i])
	}
	fmt.Println("------")
	fmt.Println("Bottom")
}
