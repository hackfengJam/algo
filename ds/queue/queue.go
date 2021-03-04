package queue

type Node interface {
}

type Queue interface {
	// 获取队列大小
	GetSize() int

	// 队列是否为空
	IsEmpty() bool

	// 元素e入队
	Enqueue(e Node)

	// 队首出队
	Dequeue() Node

	// 查看队首元素
	GetFront() Node
}
