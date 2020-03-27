// https://leetcode-cn.com/problems/climbing-stairs/submissions/

package main

func main() {

}

func climbingStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if n >= 3 {
		arr := make([]int, n+1)
		arr[0] = 1
		arr[1] = 1
		arr[2] = 2
		for i := 3; i < n+1; i++ {
			arr[i] = arr[i-1] + arr[i-2]
		}
		return arr[n]
	}
	return 0
}
