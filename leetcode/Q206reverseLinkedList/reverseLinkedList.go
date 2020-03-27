// https://leetcode-cn.com/problems/reverse-linked-list/submissions/


package main

import "fmt"

func main()  {
	var list *ListNode = &ListNode{0,nil}
	var head *ListNode
	head=list
	for i:=1;i<6;i++{
		p:=new(ListNode)
		p.Val=i
		list.Next=p
		list=list.Next
		//fmt.Println(p)
	}
	//p:=head
	p:=reverseLinkedList(head.Next)
	for p!=nil{
		fmt.Println(p.Val)
		p=p.Next
	}
}
type ListNode struct {
	Val int
	Next *ListNode
}
func reverseLinkedList(head *ListNode)  *ListNode{
	var next,pre *ListNode
	for head!=nil{
		next=head.Next
		head.Next=pre
		pre=head
		head=next
	}
	return pre
}
