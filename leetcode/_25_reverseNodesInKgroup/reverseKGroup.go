// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}
	var nodeCount int
	var newHead, pre, tmp *ListNode
	newHead = new(ListNode)
	newHead.Next = head
	pre = newHead
	for head != nil {
		nodeCount = 0
		tmp = head
		for head != nil {
			nodeCount++
			head = head.Next
			if nodeCount == k {
				break
			}
		}
		if nodeCount == k {
			var p1,pre2 *ListNode
			for nodeCount = 0; nodeCount < k; nodeCount++ {
				p1 = tmp.Next
				tmp.Next=pre2
				pre2 = tmp
				tmp = p1
			}
			pre.Next = pre2
			pre = pre.Next
			for pre!=nil && pre.Next!=nil{
				pre = pre.Next
			}
		} else {
			pre.Next = tmp
			break
		}

	}
	return newHead.Next
}
