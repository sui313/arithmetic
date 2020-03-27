// https://leetcode-cn.com/problems/3sum/submissions/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(_3Sum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(_3Sum([]int{0, 0, 0}))
}

func _3Sum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	k, i, j := 0, 0, 0
	var tmp = 0
	for ; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		if nums[k] > 0 {
			break
		}
		i, j = k+1, len(nums)-1
		for i < j {
			tmp = nums[k] + nums[i] + nums[j]
			if tmp == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				for ; i < j && nums[i] == nums[i+1]; i++ {
				} //去重
				for ; i < j && nums[j] == nums[j-1]; j-- {
				} //去重
				i++
				j--
			}
			if tmp > 0 {
				j--
			}
			if tmp < 0 {
				i++
			}
		}
	}
	return res
}
