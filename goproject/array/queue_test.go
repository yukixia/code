package array

import (
	"fmt"
	"testing"
)

/*
	循环队列相关实现
	入队
	出队
*/

type object interface{}
type CircularQueue struct {
	q        []object
	capacity int
	head     int
	tail     int
}

// 创建一个新队列
func NewCircularQueue(n int) *CircularQueue {
	if 0 == n {
		return nil
	}
	return &CircularQueue{make([]object, n), n, 0, 0}
}

// IsEmpty 队列是否为空
func (queue *CircularQueue) IsEmpty() bool {
	ret := false
	if queue.head == queue.tail {
		ret = true
	}
	return ret
}

// IsFull 队列是否满
func (queue *CircularQueue) IsFull() bool {
	ret := false
	if (queue.tail+1)%queue.capacity == queue.head {
		ret = true
	}
	return ret
}

// Enqueue 入队
func (queue *CircularQueue) Enqueue(data object) bool {
	if queue.IsFull() {
		return false
	}
	queue.q[queue.tail] = data
	queue.tail = (queue.tail + 1) % queue.capacity
	return true
}

// Dequeue 出队
func (queue *CircularQueue) Dequeue() object {
	if queue.IsEmpty() {
		return nil
	}
	data := queue.q[queue.head]
	queue.head = (queue.head + 1) % queue.capacity
	return data
}

//Print 打印
func (queue *CircularQueue) Print() {
	if queue.IsEmpty() {
		fmt.Println("queue is empty")
	}
	ret := ""
	for i := queue.head; i < queue.tail; i++ {
		ret += fmt.Sprintf("%+v", queue.q[i])
		if i+1 < queue.tail {
			ret += "=>"
		}
	}
	fmt.Printf("%s\n", ret)
}

func TestQueue(t *testing.T) {
	queue := NewCircularQueue(5)
	queue.Print()
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	fmt.Printf("%t\n", queue.Enqueue("e"))
	queue.Print()
}
