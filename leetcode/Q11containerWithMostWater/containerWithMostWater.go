// https://leetcode-cn.com/problems/container-with-most-water/submissions/


package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxA([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	max := 0
	head := 0
	tail := len(height) - 1
	area := 0
	for head != tail {

		if height[head] > height[tail] {
			area = height[tail] * (tail - head)
			tail--
		} else {
			area = height[head] * (tail - head)
			head++
		}
		if area > max {
			max = area
		}

	}
	return max
}

func maxA(height []int) int {
	max, area := 0, 0
	i, j := 0, len(height)-1
	fmt.Println(j)
	for i != j {
		if height[i] < height[j] {
			area = height[i] * (j - i)
			if max < area {
				max = area
			}
			i++
		} else {
			area = height[j] * (j - i)
			if max < area {
				max = area
			}
			j--
		}
	}
	return max
}
