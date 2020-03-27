package main

import (
	"fmt"
)

func main() {
	var list *ListNode = &ListNode{0, nil}
	var head *ListNode
	head = list
	for i := 1; i < 5; i++ {
		p := new(ListNode)
		p.Val = i
		list.Next = p
		list = list.Next
		//fmt.Println(p)
	}
	//p:=head.Next
	p := swapPairs(head.Next)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var p1, p2,v,hPrev *ListNode
	v=&ListNode{0,head}
	hPrev=v
	for head != nil && head.Next != nil {
		p1 = head //交换的第一个节点
		p2 = head.Next //交换的第二节点
		head=p2.Next  //下一次的开始节点
		// p1,p2交换位置
		p1.Next = head
		p2.Next = p1
		hPrev.Next = p2 //保存交换后的排在前面的节点
		hPrev = p1 //排在head节点前面的那个节点
	}
	return v.Next
}
