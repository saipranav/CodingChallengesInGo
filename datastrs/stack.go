package datastrs

import (
	"errors"
	"fmt"
	"strings"
)

var ErrStackUflow = errors.New("Stack Underflow while pop operation")

// Stack structure
type Stack struct {
	elements []interface{}
	top      int
}

// Creates new stack
// Syntax : stack.NewStack()
func NewStack() *Stack {
	return &Stack{}
}

// Push any number of elements to stack
func (stack *Stack) Push(value interface{}) {
	lenStackElements := len(stack.elements)

	// If elements of stack does not have sufficient capacity
	// then increase the capacity by twice
	if cap(stack.elements) <= lenStackElements+1 {
		newCap := (cap(stack.elements) + 1) * 2 // +1 if the capacity is 0
		newElements := make([]interface{}, 0, newCap)
		newElements = append(newElements, stack.elements...)
		stack.elements = newElements
	}

	// Append all values into stack elements
	stack.elements = append(stack.elements[:stack.top], value)
	stack.top = stack.top + 1
}

// Pop retrieves top element from stack and deletes the element from stack
func (stack *Stack) Pop() (interface{}, error) {

	// Handle underflow if stack is empty
	if stack.top == 0 {
		return nil, ErrStackUflow
	}
	value := stack.elements[stack.top-1]
	stack.elements[stack.top-1] = nil
	stack.top -= 1
	return value, nil
}

// Peek retrieves top element from stack
func (stack *Stack) Peek() (interface{}, error) {

	// Handle underflow if stack is empty
	if stack.top == 0 {
		return nil, ErrStackUflow
	}
	value := stack.elements[stack.top-1]
	return value, nil
}

// Empty checks if the stack is empty
func (stack *Stack) Empty() bool {

	if stack.top == 0 {
		return true
	}
	return false
}

// Convert stack to string describing the object
func (stack *Stack) String() string {
	var s string
	s += "Stack Elements :"
	var values []string
	for _, value := range stack.elements {
		values = append(values, fmt.Sprintf("%v", value))
	}
	s += " [ " + strings.Join(values, ", ") + " ] "
	s += "having top element at"
	s += " " + fmt.Sprintf("%v\n", stack.top)
	return s
}
