package queue

import (
	"fmt"
	"testing"
)

func TestLoopQueue(t *testing.T) {
	loopQueue := NewLoopQueue()
	for i := 0; i < 10; i++ {
		loopQueue.Enqueue(i)
		fmt.Printf("enqueue: %s\n", loopQueue)

		if i%3 == 2 {
			loopQueue.Dequeue()
			fmt.Printf("dequeue: %s\n", loopQueue)
		}
	}
}
