/**
* @Author: sui_liut@163.com
* @Date: 2020/3/28 21:34
 */

package main

import (
	"fmt"
)

type Node struct {
	val  int
	pre  *Node
	next *Node
}

// 创建一个双向循环链表
func insert(head *Node, node *Node) *Node {
	if head == nil {
		return node
	}
	if head.pre != nil {
		head.pre.next = node
		node.pre = head.pre
		node.next = head
		head.pre = node
	} else {
		head.pre = node
		head.next = node
		node.pre = head
		node.next = head
	}
	return head
}

func main() {
	head := &Node{0, nil, nil}
	for i := 1; i < 10; i++ {
		insert(head, &Node{i, nil, nil})
	}
	p := head.next
	fmt.Println(head.val)
	for p != head {
		fmt.Println(p.val)
		p = p.next
	}

	p = p.pre
	for p != head {
		fmt.Println(p.val)
		p = p.pre
	}
}
