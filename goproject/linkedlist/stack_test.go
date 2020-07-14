/*
	stack基本概念
	先进后出
*/

package linkedlist

import (
	"fmt"
	"testing"
)

type object interface{}

type node struct {
	val  object
	next *node
}

type StackList struct {
	topNode *node
}

func NewStackList() *StackList {
	return &StackList{topNode: nil}
}

// IsEmpty 检测栈是否为空
func (stack *StackList) IsEmpty() bool {
	return stack.topNode == nil
}

// Push 入栈
func (stack *StackList) Push(data object) {
	temp := &node{
		val:  data,
		next: stack.topNode,
	}
	stack.topNode = temp
}

// Pop 出栈
func (stack *StackList) Pop() object {
	if true == stack.IsEmpty() {
		return nil
	}
	temp := stack.topNode.next
	curNode := stack.topNode
	stack.topNode = temp
	return curNode.val
}

// Top 获取栈顶元素
func (stack *StackList) Top() object {
	if true == stack.IsEmpty() {
		return nil
	}
	return stack.topNode.val
}

// Flush 刷新栈
func (stack *StackList) Flush() {
	stack.topNode = nil
}

// 打印栈元素
func (stack *StackList) Print() {
	curNode := stack.topNode
	format := ""
	for nil != curNode {
		format = format + fmt.Sprintf("%+v", curNode.val)
		curNode = curNode.next
		if nil != curNode {
			format = format + "=>"
		}
	}
	fmt.Println(format)
}

//
func TestStact(t *testing.T) {
	stacktest := NewStackList()
	fmt.Printf("%t\n", stacktest.IsEmpty())
	stacktest.Push(1)
	stacktest.Push(2)
	stacktest.Push(3)
	stacktest.Push(4)
	fmt.Printf("%d\n", stacktest.Top())
	stacktest.Print()
}
