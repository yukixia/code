package linkedlist

import (
	"fmt"
	"testing"
)

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
	length   int
}

func NewNode(data Object) *Node {
	return &Node{Data: data}
}

func NewList() *List {
	return &List{headNode: NewNode(0), length: 0}
}

func (node *Node) GetValue() Object {
	return node.Data
}

// Empty 链表是否位空
func (list *List) Empty() bool {
	ret := true
	if 0 == list.length {
		ret = false
	}
	return ret
}

// Length 链表的长度
func (list *List) Length() int {
	return list.length
}

// InsertAfter 在某个节点之后插入
func (list *List) InsertAfter(node *Node, data Object) bool {
	ret := true
	if nil == node {
		ret = false
	} else {
		temp := NewNode(data)
		temp.Next = node.Next
		node.Next = temp
		list.length++
	}
	return ret
}

// InsertBefore 在某个节点之前插入
func (list *List) InsertBefore(node *Node, data Object) bool {
	ret := true
	if nil == node {
		ret = false
	} else {
		preNode := list.headNode
		cur := preNode.Next
		for nil != cur {
			if node == cur {
				break
			}
			preNode = cur
			cur = cur.Next
		}
		if nil == cur {
			ret = false
		} else {
			temp := NewNode(data)
			preNode.Next = temp
			temp.Next = cur
			list.length++
		}
	}
	return ret
}

// InsrtToHead 在头部插入
func (list *List) InsertToHead(data Object) bool {
	return list.InsertAfter(list.headNode, data)
}

// InsertToTail 在尾部插入
func (list *List) InsertToTail(data Object) bool {
	cur := list.headNode
	for nil != cur.Next {
		cur = cur.Next
	}
	return list.InsertAfter(cur, data)
}

// Delete 删除某个节点
func (list *List) DeleteNode(node *Node) bool {
	ret := true
	if nil == node {
		ret = false
	} else {
		preNode := list.headNode
		curNode := preNode.Next
		for nil != curNode {
			if node == curNode {
				break
			}
			preNode = curNode
			curNode = curNode.Next
		}
		if nil != curNode {
			preNode.Next = curNode.Next
			curNode = nil
			list.length--
		} else {
			ret = false
		}
	}
	return ret
}

// 删除链表中的某个值
func (list *List) DeleteValue(data Object) bool {
	curNode := list.headNode.Next
	for nil != curNode {
		if data == curNode.GetValue() {
			break
		}
		curNode = curNode.Next
	}
	return list.DeleteNode(curNode)
}

// FindByValue 通过值查找对应节点
func (list *List) FindByValue(value Object) *Node {
	curNode := list.headNode.Next
	for nil != curNode {
		if value == curNode.GetValue() {
			break
		}
		curNode = curNode.Next
	}
	return curNode
}

// ReverseList 链表反转
func (list *List) ReverseList() *Node {
	if 0 == list.length || 1 == list.length {
		return list.headNode
	}
	var preNode *Node = nil
	curNode := list.headNode.Next
	for nil != curNode.Next {
		temp := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = temp
	}
	//temp := curNode.Next
	curNode.Next = preNode
	//preNode.Next = temp
	list.headNode.Next = curNode
	return list.headNode
}

// isPalindrome2 判断是否为回文链表
func (list *List) isPalindrome2() bool {
	isPalindrome2 := true
	if 0 == list.length || 1 == list.length {
		return isPalindrome2
	}
	var preNode *Node = nil
	len := list.length
	curNode := list.headNode.Next
	step := len / 2
	//reverse
	for i := 1; i <= step; i++ {
		temp := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = temp
	}
	mid := curNode
	var left, right *Node = preNode, nil
	if len%2 == 0 {
		right = mid
	} else {
		right = mid.Next
	}
	//compare
	for nil != left && nil != right {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome2 = false
			break
		}
		left = left.Next
		right = right.Next
	}
	//reduction
	curNode = preNode
	preNode = mid
	for nil != curNode {
		temp := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = temp
	}
	list.headNode.Next = preNode
	return isPalindrome2
}

// Print 打印链表
func (list *List) Print() {
	cur := list.headNode.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.Next
		if nil != cur {
			format += "=>"
		}
	}
	fmt.Println(format)
}

func TestList2(t *testing.T) {
	listtest := NewList()
	listtest.InsertToHead(1)
	listtest.InsertToHead(2)
	listtest.InsertToHead(3)
	listtest.InsertToTail('a')
	listtest.InsertToTail('b')
	listtest.InsertToTail('c')
	listtest.Print()
	fmt.Println(listtest.Length())
	listtest.DeleteValue('c')
	listtest.Print()
	fmt.Println(listtest.Length())
	node := listtest.FindByValue('a')
	fmt.Printf("+%v\n", node)
	listtest.ReverseList()
	listtest.Print()
	listtest.ReverseList()
	listtest.Print()
	listtest2 := NewList()
	listtest2.InsertToTail("a")
	listtest2.InsertToTail("b")
	listtest2.InsertToTail("c")
	listtest2.InsertToTail("d")
	listtest2.InsertToTail("c")
	listtest2.InsertToTail("b")
	listtest2.InsertToTail("a")
	listtest2.Print()
	fmt.Printf("%t\n", listtest2.isPalindrome2())
	listtest2.Print()
}
