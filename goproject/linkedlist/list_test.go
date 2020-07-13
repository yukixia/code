package linkedlist

import (
	"fmt"
	"testing"
)

// Node 定义链表结构体
type Node struct {
	Val  int
	Next *Node
}

// Shownode 遍历列表
func Shownode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}

// Insertfrontnode 头部插入节点
func Insertfrontnode(head *Node, val int) *Node {
	Shownode(head)
	node := &Node{
		Val: val,
	}
	temp := head
	head = node
	node.Next = temp
	Shownode(head)
	return head
}

// Inserttailnode 尾部插入节点
func Inserttailnode(head *Node, val int) *Node {
	// node := &Node{
	// 	Val: val,
	// }
	node := new(Node)
	node.Val = val
	p := head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
	return head
}

func TestList(t *testing.T) {
	head := &Node{Val: 0}
	Insertfrontnode(head, 1)
	Shownode(head)

}
