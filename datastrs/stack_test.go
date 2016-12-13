package datastrs

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	testingStack := New()
	fmt.Print(testingStack.String())

	// Output:
	// [] 0
}

func TestStack_PushIntegers(t *testing.T) {
	testingStack := New()
	testingStack.Push(1)
	testingStack.Push(2)
	testingStack.Push(3)
	testingStack.Push(4, 5, 6, 7, 8)
	fmt.Print(testingStack.String())

	// Output:
	// [ 1, 2, 3, 4, 5, 6, 7, 8 ] having top element at 8
}

func TestStack_PushStrings(t *testing.T) {
	testingStack := New()
	testingStack.Push("a")
	testingStack.Push("b")
	testingStack.Push("c")
	testingStack.Push("d", "ef", "g")
	fmt.Print(testingStack.String())

	// Output:
	// [ a, b, c, d, ef, g ] having top element at 6
}

func TestStack_Pop(t *testing.T) {
	testingStack := New()
	testingStack.Push("a")
	testingStack.Push("b")
	value, error := testingStack.Pop()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Print("Popped value -> " + fmt.Sprintf("%v\n", value))
	fmt.Print(testingStack.String())

	// Output:
	// Popped value ->  b
	// [ a, <nil> ] having top element at 1
}

func TestStack_PopEmptyStack(t *testing.T) {
	testingStack := New()
	_, error := testingStack.Pop()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Print(testingStack.String())

	// Output:
	// Stack Underflow while pop operation
	// Stack Elements : [  ] having top element at 0
}

func TestStack_PopAndPush(t *testing.T) {
	testingStack := New()
	testingStack.Push("a")
	testingStack.Push(2)
	value, error := testingStack.Pop()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Print("Popped value -> " + fmt.Sprintf("%v\n", value))
	fmt.Print(testingStack.String())
	testingStack.Push("b")
	fmt.Print(testingStack.String())

	// Output:
	// Popped value -> 2
	// Stack Elements : [ a, <nil> ] having top element at 1
	// Stack Elements : [ a, b ] having top element at 2
}

func TestStack_Peek(t *testing.T) {
	testingStack := New()
	testingStack.Push("a")
	testingStack.Push(2)
	value, error := testingStack.Peek()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Print("Peeked value -> " + fmt.Sprintf("%v\n", value))
	fmt.Print(testingStack.String())

	// Output:
	// Peeked value -> 2
	// Stack Elements : [ a, 2 ] having top element at 2
}

func TestStack_PeekEmptyStack(t *testing.T) {
	testingStack := New()
	_, error := testingStack.Peek()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Print(testingStack.String())

	// Output:
	// Stack Underflow while pop operation
	// Stack Elements : [  ] having top element at 0
}
