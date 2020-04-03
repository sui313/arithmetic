/**
* @Author: sui_liut@163.com
* @Date: 2020/3/28 1:22
 */

package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Rigth *Node
}

type Result struct {
	MaxDepth    int
	MaxDistance int
}

var p1 = &Node{1, p2, p3}
var p2 = &Node{2, p4, p5}
var p3 = &Node{3, p6, p7}
var p4 = &Node{4, p8, nil}
var p5 = &Node{5, nil, nil}
var p6 = &Node{6, nil, nil}
var p7 = &Node{7, p9, nil}
var p8 = &Node{8, nil, nil}
var p9 = &Node{9, nil, nil}

var n1 = &Node{1, n2, n3}
var n2 = &Node{2, n4, n5}
var n3 = &Node{3, nil, nil}
var n4 = &Node{4, n6, n7}
var n5 = &Node{5, n8, nil}
var n6 = &Node{6, nil, nil}
var n7 = &Node{7, nil, n9}
var n8 = &Node{8, nil, nil}
var n9 = &Node{9, nil, nil}

// 二叉树中最远两个节点的距离
func findMaxDistance(p *Node) *Result {
	var res = new(Result)
	if p == nil {
		res.MaxDepth = -1
		res.MaxDistance = -1
		return res
	}

	l := findMaxDistance(p.Left)
	r := findMaxDistance(p.Rigth)

	res.MaxDepth = max(l.MaxDepth, r.MaxDepth) + 1

	res.MaxDistance = max(max(l.MaxDistance, r.MaxDistance), l.MaxDepth+r.MaxDepth+2)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s1 := findMaxDistance(p1)
	fmt.Println(s1)
	fmt.Println(findMaxDistance(n1))
}
