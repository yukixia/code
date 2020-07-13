package tree

import "fmt"

// Node struct
type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
	Num   int
}

func compareto(key1 int, key2 int) int {
	if key1 < key2 {
		return -1
	} else if key1 > key2 {
		return 1
	}
	return 0
}

// Size 获取二叉树的大小
func Size(root *Node) int {
	if nil == root {
		return 0
	}
	return root.Num
}

// Get 获取key对应的值
func Get(root *Node, key int) string {
	p := root
	for nil != p {
		cmp := compareto(key, p.Key)
		if 0 == cmp {
			return p.Value
		} else if 0 > cmp {
			p = p.Left
		} else if 0 < cmp {
			p = p.Right
		}
	}
	return ""
}

// Put 插入元素
func Put(root *Node, key int, val string) *Node {
	p := root
	if nil == p {
		return &Node{
			Key:   key,
			Value: val,
			Num:   1,
		}
	}
	for nil != p {
		cmp := compareto(key, p.Key)
		if 0 == cmp {
			p.Value = val
			return root
		} else if 0 > cmp {
			if nil == p.Left {
				node := new(Node)
				node.Key = key
				node.Value = val
				node.Num = 1
				p.Left = node
				p.Num = Size(p.Left) + Size(p.Right) + 1
				break
			} else {
				p = p.Left
			}
		} else if 0 < cmp {
			if nil == p.Right {
				node := new(Node)
				node.Key = key
				node.Value = val
				node.Num = 1
				p.Right = node
				p.Num = Size(p.Left) + Size(p.Right) + 1
				break
			} else {
				p = p.Right
			}
		}
	}
	return root

}

// Inorder 中序遍历
func Inorder(root *Node) {
	p := root
	if nil == p {
		return
	}
	Inorder(p.Left)
	fmt.Printf("%s,", p.Value)
	Inorder(p.Right)
}

// Printlevel 层遍历
func Printlevel(root *Node) {
	p := root
	var slicequeue []*Node
	if nil != p {
		slicequeue = append(slicequeue, p)
	}
	for len(slicequeue) != 0 {
		p = slicequeue[0]
		slicequeue = slicequeue[1:]
		fmt.Println(p.Value)
		if p.Left != nil {
			slicequeue = append(slicequeue, p.Left)
		}
		if p.Right != nil {
			slicequeue = append(slicequeue, p.Right)
		}
	}
}

// Rank 返回key在二叉树中的排名
// 如果key等于根结点，排名为size(root.left)
// 如果key小于根结点的值，返回在左子树的排名
// 如果key大于根结点，返回size(root.left)+1+在右子树的排名
func Rank(root *Node, key int) int {
	p := root
	rank := 0
	for nil != p {
		cmp := compareto(key, p.Key)
		if 0 == cmp {
			rank = Size(p.Left)
		} else if 0 > cmp {
			p = p.Left
		} else if 0 < cmp {
			rank = rank + Size(p.Left) + 1
			p = p.Right
		}
	}
	return rank
}
