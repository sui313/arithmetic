package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode, pos int) *ListNode {

	if head == nil || head.Next == nil || pos < 0 {
		return nil
	}

	var ptr1, ptr2, slow, fast *ListNode

	slow = head
	fast = head.Next
	ptr1 = head

	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			ptr2 = slow.Next
			for ptr1 != ptr2 {
				ptr1 = ptr1.Next
				ptr2 = ptr2.Next
			}
			return ptr1
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
