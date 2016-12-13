package datastrs

import (
	"errors"
	"fmt"
	"strings"
)

var ErrUnderflow = errors.New("Stack Underflow while pop operation")

// Stack structure
type Stack struct {

	// Elements of the stack are stored in array
	elements []interface{}

	// Stack's top index
	top int
}

// New function creates new stack
// Syntax : stack.New()
func New() *Stack {
	return &Stack{}
}

// Push any number of elements to stack
func (stack *Stack) Push(values ...interface{}) {
	lenStackElements := len(stack.elements)
	lenValues := len(values)

	// If elements of stack does not have sufficient capacity
	// then increase the capacity by twice
	if cap(stack.elements) <= lenStackElements+lenValues {
		newCap := (cap(stack.elements) + lenValues + 1) * 2 // +1 if the capacity is 0
		newElements := make([]interface{}, 0, newCap)
		newElements = append(newElements, stack.elements...)
		stack.elements = newElements
	}

	// Append all values into stack elements
	stack.elements = append(stack.elements[:stack.top], values...)
	stack.top = stack.top + lenValues
}

// Pop retrieves top element from stack and deletes the element from stack
func (stack *Stack) Pop() (interface{}, error) {

	// Handle underflow if stack is empty
	if stack.top == 0 {
		return nil, ErrUnderflow
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
		return nil, ErrUnderflow
	}
	value := stack.elements[stack.top-1]
	return value, nil
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
