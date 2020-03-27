/*
归纳：

   对于螺旋矩阵之类的题目有很多变形，但是除了边界条件不同外，都是两层循环的左右边界向外扩展或者向内收敛的问题，因此一般只要控制好循环左右上下边界的增减，问题便解决了

给定一个包含 m x n 个元素的矩阵（m 行, n列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]


给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/


package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	fmt.Println(spiral(arr))
	arr2 := spiral2(3)
	for _, v := range arr2 {
		fmt.Println(v)
	}
}

func spiral(arr [][]int) []int {
	if len(arr) == 0 {
		return nil
	}

	up := 0
	down := len(arr)
	left := 0
	right := len(arr[0])
	result := make([]int, 0, down*right)
	for up < down && left < right {
		up1 := left
		for up1 < right {
			result = append(result, arr[left][up1])
			up1++
		}
		up++

		right1 := up
		for right1 < down {
			result = append(result, arr[right1][right-1])
			right1++
		}
		right = right - 1

		down1 := right - 1
		for down1 >= left {
			result = append(result, arr[down-1][down1])
			down1--
		}
		down = down - 1
		left1 := down - 1
		for left1 >= up {
			result = append(result, arr[left1][left])
			left1--
		}
		left++
	}
	return result
}

func spiral2(num int) [][]int {

	up := 0
	down := num
	left := 0
	right := num
	result := make([][]int, num)
	for i := 0; i < num; i++ {
		result[i] = make([]int, num)
	}
	k := 1
	for up < down && left < right && k <= num*num {
		up1 := left
		for up1 < right {
			result[left][up1] = k
			k++
			up1++
		}
		up++

		right1 := up
		for right1 < down {
			result[right1][right-1] = k
			k++
			right1++
		}
		right = right - 1

		down1 := right - 1
		for down1 >= left {
			result[down-1][down1] = k
			k++
			down1--
		}

		down = down - 1
		left1 := down - 1
		for left1 >= up {
			result[left1][left] = k
			k++
			left1--
		}
		left++
	}
	return result
}
