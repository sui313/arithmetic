// https://leetcode-cn.com/problems/move-zeroes/
package main

import "fmt"

func main() {
	var nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if j != i {
				nums[i] = 0
			}
			j++
		}
	}
}
