// https://leetcode-cn.com/problems/linked-list-cycle/
// 找列表中的环
package main

import (
	"fmt"
	"time"
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
	//list.Next = head.Next
	fmt.Println(hasCycle(head))
	fmt.Println(detectCycle(head, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var quik, slow *ListNode
	slow = head.Next
	quik = head.Next.Next
	for slow != nil && slow.Next != nil &&
		quik != nil && quik.Next != nil {
		if slow == quik {
			return true
		}
		slow = slow.Next
		quik = quik.Next.Next
	}
	return false
}

func detectCycle(head *ListNode, pos int) *ListNode {
	if head == nil || head.Next == nil || pos < 0 {
		return nil
	}

	var quik, slow *ListNode
	slow = head

	n := 0
	for slow != nil {// 构造环
		time.Sleep(time.Second)
		fmt.Println(slow)
		if slow.Next == nil {
			slow.Next = quik
			break
		}
		if n == pos {// 环入口点
			quik = slow
		}
		n++
		slow = slow.Next
	}

	slow = head
	quik = head.Next
	var p1, p2 *ListNode
	p1 = head
	for slow != nil && slow.Next != nil &&
		quik != nil && quik.Next != nil {
		if slow == quik {
			p2 = slow.Next
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}
			fmt.Println("11111", p1, p2)
			return p1
		}
		slow = slow.Next
		quik = quik.Next.Next
	}
	return nil
}
