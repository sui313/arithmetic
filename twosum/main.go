package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	demo2()
}
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func demo1() {

	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	a, b := make([]int, 100), make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = int(ra.Int31n(100) - 20)
	}
	for i := 0; i < 100; i++ {
		b[i] = int(ra.Int31n(100) - 50)
	}
	var s1, s2 int
	for i := 0; i < 100; i++ {
		s1 = a[i] + b[i]
		s2 = sum(a[i], b[i])
		if s1 != s2 {
			fmt.Println("错误:", a[i], b[i])
			return
		} else {
			fmt.Println(i, "::", s1, "=", a[i], "+", b[i])
		}
	}
}

func sum(a, b int) int {
	yh := XOR(a, b)
	t := yh
	y := AND(a, b)
	//fmt.Printf("A:%d:%08b\n%d:%08b\n\n", yh, yh, y, y)
	for y != 0 {
		y = y << 1
		//fmt.Printf("B:%d:%08b\n%d:%08b\n\n", yh, yh, y, y)
		yh = XOR(yh, y)
		//fmt.Printf("C:%d:%08b\n%d:%08b\n\n", yh, yh, y, y)
		y = AND(t, y)
		//fmt.Printf("D:%d:%08b\n%d:%08b\n\n", yh, yh, y, y)
	}
	//fmt.Println(yh)
	return yh
}

func XOR(a, b int) int {
	return a ^ b
}

func AND(a, b int) int {
	return a & b
}

func demo2() {
	NUM:=1
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	a, b := make([]int, NUM), make([]int, NUM)
	for i := 0; i < NUM; i++ {
		a[i] = int(ra.Int31n(10))
		if i == NUM-1 && a[NUM-1] == 0 {
			i--
		}
	}
	for i := 0; i < NUM; i++ {
		b[i] = int(ra.Int31n(10))
		if i == NUM-1 && b[NUM-1] == 0 {
			i--
		}
	}
	fmt.Println(a, b)
	l1 := createList(a)
	l2 := createList(b)
	//for {
	//	if l1 == nil {
	//		break
	//	}
	//	fmt.Printf("%d_", l1.Val)
	//	l1=l1.Next
	//}
	//fmt.Println()
	//for {
	//	if l2 == nil {
	//		break
	//	}
	//	fmt.Printf("%d_", l2.Val)
	//	l2=l2.Next
	//}
	//fmt.Println()
	start := time.Now()

	l3 := addTwoNumbers(l1, l2)
	end:=time.Now()
	for {
		if l3 == nil {
			break
		}
		fmt.Printf("%d", l3.Val)
		l3 = l3.Next
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(end.UnixNano()-start.UnixNano())
}
/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(arr []int) *ListNode {
	var p *ListNode
	var l = new(ListNode)
	p = l
	for _, v := range arr {
		l.Next = new(ListNode)
		l.Next.Val = v
		l = l.Next
	}
	return p.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var l3 = new(ListNode)
	var p = l3
	var n int
	for {
		if l1 != nil {
			l3.Val = l3.Val + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l3.Val = l3.Val + l2.Val
			l2 = l2.Next
		}
		n = l3.Val / 10
		l3.Val = l3.Val % 10
		if n > 0 {
			l3.Next = new(ListNode)
			l3.Next.Val = n
			n = 0
		}

		if l1 == nil && l2 == nil {
			break
		}

		if l3.Next == nil {
			l3.Next = new(ListNode)
		}
		l3 = l3.Next
	}
	return p
}
