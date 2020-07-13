package leecode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	reminder := 0
	quatient := 0
	head := new(ListNode)
	curNode := head
	lastNode := curNode
	for nil != l1 && nil != l2 {
		result := l1.Val + l2.Val + quatient
		reminder = result % 10
		quatient = result / 10
		curNode.Val = reminder
		tmp := new(ListNode)
		curNode.Next = tmp
		lastNode = curNode
		curNode = curNode.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for nil != l1 {
		result := l1.Val + quatient
		reminder = result % 10
		quatient = result / 10
		curNode.Val = reminder
		tmp := new(ListNode)
		curNode.Next = tmp
		lastNode = curNode
		curNode = curNode.Next
		l1 = l1.Next
	}
	for nil != l2 {
		result := l2.Val + quatient
		reminder = result % 10
		quatient = result / 10
		curNode.Val = reminder
		tmp := new(ListNode)
		curNode.Next = tmp
		lastNode = curNode
		curNode = curNode.Next
		l2 = l2.Next
	}
	if 0 != quatient {
		curNode.Val = quatient
		curNode.Next = nil
		lastNode = curNode
	}
	lastNode.Next = nil
	return head
}
