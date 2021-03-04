package queue

// source: https://go2goplay.golang.org/p/Y9M55pYaP4R
/*
package main

import (
	"fmt"
	"strings"
)

// The playground now uses square brackets for type parameters. Otherwise,
// the syntax of type parameter lists matches the one of regular parameter
// lists except that all type parameters must have a name, and the type
// parameter list cannot be empty. The predeclared identifier "any" may be
// used in the position of a type parameter constraint (and only there);
// it indicates that there are no constraints.

// 假设 a wrapper class which for int64
type Int64 struct {
	i int
}

// Queue
type Queue[T any] interface {
    // 获取队列大小
    GetSize() int

    // 队列是否为空
    IsEmpty() bool

    // 元素e入队
    Enqueue(e *T)

    // 队首出队
    Dequeue() *T

    // 查看队首元素
    GetFront() *T
}

// 循环队列
type LoopQueue[T any] struct {
	data        []*T
	front, tail int
	size        int
}

func NewLoopQueueWithCapacity[T any](capacity int) *LoopQueue[T] {
	l := &LoopQueue[T]{
		data:  make([]*T, capacity+1),
		front: 0,
		tail:  0,
		size:  0,
	}
	return l
}

func NewLoopQueue[T any]() *LoopQueue[T] {
	return NewLoopQueueWithCapacity[T](10)
}

func (l *LoopQueue[T]) GetSize() int {
	return l.size
}

func (l *LoopQueue[T]) getCapacity() int {
	return len(l.data) - 1
}

func (l *LoopQueue[T]) resize(newCapacity int) {
	newData := make([]*T, newCapacity+1)

	ln := len(l.data)
	for i := 0; i < l.size; i++ {
		newData[i] = l.data[(i+l.front)%ln]
	}
	l.data = newData
	l.front = 0
	l.tail = l.size
	return
}

func (l *LoopQueue[T]) IsEmpty() bool {
	return l.front == l.tail
}

func (l *LoopQueue[T]) Enqueue(e *T) {
	if (l.tail+1)%len(l.data) == l.front {
		l.resize(l.getCapacity() * 2)
	}

	l.data[l.tail] = e
	l.tail = (l.tail + 1) % len(l.data)
	l.size++
}

func (l *LoopQueue[T]) Dequeue() *T {
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

func (l *LoopQueue[T]) GetFront() *T {
	if l.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	return l.data[l.front]
}

func (l *LoopQueue[T]) String() string {
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

func main() {
	var loopQueue Queue[Int64]
	loopQueue = NewLoopQueue[Int64]()
	for i := 0; i < 10; i++ {
		loopQueue.Enqueue(&Int64{i: i})
		fmt.Printf("enqueue: %s\n", loopQueue)

		if i%3 == 2 {
			loopQueue.Dequeue()
			fmt.Printf("dequeue: %s\n", loopQueue)
		}
	}
}
*/
