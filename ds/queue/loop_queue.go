package queue

import (
	"fmt"
	"strings"
)

type LoopQueue struct {
	data        []Node
	front, tail int
	size        int
}

func NewLoopQueue() *LoopQueue {
	return NewLoopQueueWithCapacity(10)
}

func NewLoopQueueWithCapacity(capacity int) *LoopQueue {
	l := &LoopQueue{
		data:  make([]Node, capacity+1),
		front: 0,
		tail:  0,
		size:  0,
	}
	return l
}

func (l *LoopQueue) GetSize() int {
	return l.size
}

func (l *LoopQueue) getCapacity() int {
	return len(l.data) - 1
}

func (l *LoopQueue) resize(newCapacity int) {
	newData := make([]Node, newCapacity+1)

	ln := len(l.data)
	for i := 0; i < l.size; i++ {
		newData[i] = l.data[(i+l.front)%ln]
	}
	l.data = newData
	l.front = 0
	l.tail = l.size
	return
}

func (l *LoopQueue) IsEmpty() bool {
	return l.front == l.tail
}

func (l *LoopQueue) Enqueue(e Node) {
	if (l.tail+1)%len(l.data) == l.front {
		l.resize(l.getCapacity() * 2)
	}

	l.data[l.tail] = e
	l.tail = (l.tail + 1) % len(l.data)
	l.size++
}

func (l *LoopQueue) Dequeue() Node {
	if l.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}

	ret := l.data[l.front]
	l.data[l.front] = nil
	l.front = (l.front + 1) % len(l.data)
	l.size--

	if l.size == l.getCapacity()/4 && l.getCapacity()/2 != 0 {
		l.resize(l.getCapacity() / 2)
	}
	return ret
}

func (l *LoopQueue) GetFront() Node {
	if l.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	return l.data[l.front]
}

func (l *LoopQueue) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", l.size, l.getCapacity()))
	sb.WriteString("front [")

	ln := len(l.data)
	for i := l.front; i != l.tail; i = (i + 1) % ln {
		sb.WriteString(fmt.Sprintf("%v", l.data[i]))
		if (i+1)%ln != l.tail {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("] tail")
	return sb.String()
}

