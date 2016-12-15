package datastrs

import (
	"errors"
	"fmt"
	"strings"
)

var ErrQOflow = errors.New("Queue Overflow while enqueue operation")
var ErrQUflow = errors.New("Queue Underflow while dequeue operation")

// Queue structure
type Queue struct {
	elements []interface{}
	head     int
	tail     int
}

// Creates new queue with specific size
// Syntax : queue.NewQueue(size)
func NewQueue(size int) *Queue {
	intialQueue := make([]interface{}, size, size)
	return &Queue{intialQueue, 0, 0}
}

// Enqueue the element if not possible throw overflow error
func (queue *Queue) Enqueue(value interface{}) (bool, error) {
	// If tail has reached the end of list
	if queue.tail == len(queue.elements) {
		// If tail has seen the head we have reached full capacity in the queue
		// Reset tail to point from the start of queue
		queue.tail = 0

		if queue.tail == queue.head {
			// If queue is stagnant then reset queue tail to last position + 1
			// so that execution checks the above top if every time invoked
			queue.tail = len(queue.elements)
			return false, ErrQOflow
		}
	}
	queue.elements[queue.tail] = value
	queue.tail = queue.tail + 1
	return true, nil
}

// Dequeue the element if not possible throw overflow error
func (queue *Queue) Dequeue() (interface{}, error) {
	// If head has reached end of list
	if queue.head == len(queue.elements) {
		// If head has seen tail we have completed work for all queue elements
		if queue.head == queue.tail {
			queue.head = len(queue.elements)
			return false, ErrQUflow
		}
		queue.head = 0
	}
	value := queue.elements[queue.head]
	queue.elements[queue.head] = nil
	queue.head = queue.head + 1
	return value, nil
}

// Queue is empty if head and tail points to same position
func (queue *Queue) Empty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

// Convert stack to string describing the object
func (queue *Queue) String() string {
	var s string
	s += "Queue Elements :"
	var values []string
	for _, value := range queue.elements {
		values = append(values, fmt.Sprintf("%v", value))
	}
	s += " [ " + strings.Join(values, ", ") + " ] "
	s += "having head element at"
	s += " " + fmt.Sprintf("%v ", queue.head)
	s += "having tail element at"
	s += " " + fmt.Sprintf("%v\n", queue.tail)
	return s
}
