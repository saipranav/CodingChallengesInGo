package datastrs

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	testingQueue := NewQueue(10)
	fmt.Print(testingQueue.String())

	// Output:
	// Queue Elements : [ <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil> ] having head element at 0 having tail element at 0
}

func TestQueue_Enqueue(t *testing.T) {
	testingQueue := NewQueue(10)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(1)
	fmt.Print(testingQueue.String())

	// Output:
	// Queue Elements : [ 1, 1, 1, 1, <nil>, <nil>, <nil>, <nil>, <nil>, <nil> ] having head element at 0 having tail element at 4
}

func TestQueue_EnqueueOverflow(t *testing.T) {
	testingQueue := NewQueue(10)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(2)
	testingQueue.Enqueue(3)
	testingQueue.Enqueue(4)
	testingQueue.Enqueue(5)
	testingQueue.Enqueue(6)
	testingQueue.Enqueue(7)
	testingQueue.Enqueue(8)
	testingQueue.Enqueue(9)
	_, error := testingQueue.Enqueue(10)
	if error != nil {
		fmt.Println(error)
	}
	_, error = testingQueue.Enqueue(11)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Print(testingQueue.String())

	// Output:
	// Queue Overflow while enqueue operation
	// Queue Elements : [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ] having head element at 0 having tail element at 10
}

func TestQueue_Dequeue(t *testing.T) {
	testingQueue := NewQueue(10)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(1)
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	fmt.Print(testingQueue.String())

	// Output:
	// Queue Elements : [ <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil> ] having head element at 2 having tail element at 2
}

func TestQueue_DequeueOverflow(t *testing.T) {
	testingQueue := NewQueue(10)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(2)
	testingQueue.Enqueue(3)
	testingQueue.Enqueue(4)
	testingQueue.Enqueue(5)
	testingQueue.Enqueue(6)
	testingQueue.Enqueue(7)
	testingQueue.Enqueue(8)
	testingQueue.Enqueue(9)
	testingQueue.Enqueue(10)
	fmt.Print(testingQueue.String())

	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	value, error := testingQueue.Dequeue()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(value)
	_, error = testingQueue.Dequeue()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Print(testingQueue.String())

	// Output:
	// Queue Underflow while dequeue operation
	// Queue Elements : [ <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil> ] having head element at 10 having tail element at 10
}

func TestQueue_EnqueueDequeue(t *testing.T) {
	testingQueue := NewQueue(10)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(2)
	testingQueue.Enqueue(3)
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Enqueue(4)
	testingQueue.Enqueue(5)
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Enqueue(6)
	testingQueue.Enqueue(7)
	testingQueue.Enqueue(8)
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Enqueue(9)
	testingQueue.Enqueue(10)
	testingQueue.Dequeue()
	fmt.Print(testingQueue.String())

	// Output:
	// Queue Elements : [ <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, <nil>, 10 ] having head element at 9 having tail element at 10
}

func TestQueue_EnqueueDequeueAfterCapacity(t *testing.T) {
	testingQueue := NewQueue(10)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(2)
	testingQueue.Enqueue(3)
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Enqueue(4)
	testingQueue.Enqueue(5)
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Enqueue(6)
	testingQueue.Enqueue(7)
	testingQueue.Enqueue(8)
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Enqueue(9)
	testingQueue.Enqueue(10)
	testingQueue.Enqueue(1)
	testingQueue.Enqueue(2)
	testingQueue.Enqueue(3)
	testingQueue.Enqueue(4)
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	testingQueue.Dequeue()
	fmt.Print(testingQueue.String())

	// Output:
	// Queue Elements : [ <nil>, 2, 3, 4, <nil>, <nil>, <nil>, <nil>, <nil>, <nil> ] having head element at 1 having tail element at 4
}
