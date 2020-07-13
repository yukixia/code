package main

import (
	"fmt"
	"goproject/tree"
)

func main() {
	// head := new(linkedlist.Node)
	// head.Val = 1
	// node1 := new(linkedlist.Node)
	// node1.Val = 2
	// head.Next = node1
	// linkedlist.Shownode(head)
	// head = linkedlist.Insertfrontnode(head, 3)
	// linkedlist.Shownode(head)
	// head = linkedlist.Inserttailnode(head, 4)
	// linkedlist.Shownode(head)

	//list := [][]int{{1, 4}, {0, 4}}
	//list := []int{4, 3, 1, 5, 6}
	//sortalgorithm.UBSort(list)
	//fmt.Println(list)
	head := tree.Put(nil, 4, "4")
	head = tree.Put(head, 2, "2")
	head = tree.Put(head, 3, "3")
	head = tree.Put(head, 7, "7")
	head = tree.Put(head, 5, "5")
	//val := tree.Get(head, 4)
	//fmt.Println(val)
	tree.Inorder(head)
	fmt.Println()
	tree.Printlevel(head)
}
