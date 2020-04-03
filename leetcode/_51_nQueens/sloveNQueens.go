/**
* @Author: sui_liut@163.com
* @Date: 2020/3/30 12:27
 */

package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	slove := make([][]int, 0)
	pie := make(map[int]bool, 2*n-1)
	na := make(map[int]bool, 2*n-1)
	state := make([]int, 0, n)
	dfs(n, 0, state, pie, na, &slove)
	var result [][]string
	for _, v := range slove {
		var res []string
		for row := 0; row < n; row++ {
			s := ""
			for col := 0; col < n; col++ {
				if col == v[row] {
					s = s + "Q"
				} else {
					s = s + "."
				}
			}
			res = append(res, s)
		}
		result = append(result, res)
	}
	return result
}

func dfs(n, row int, state []int, pie, na map[int]bool, slove *[][]int) {
	if row >= n {
		arr := make([]int, n)
		copy(arr, state)
		*slove = append(*slove, arr)
		return
	}

	for col := 0; col < n; col++ {
		if !isvalied(row, col, state, pie, na) {
			continue
		}
		state = append(state, col)
		pie[row+col] = true
		na[row-col] = true
		dfs(n, row+1, state, pie, na, slove)
		state = state[0 : len(state)-1]
		delete(pie, row+col)
		delete(na, row-col)
	}
}

func isvalied(row, col int, lie []int, pie, na map[int]bool) bool {
	for _, v := range lie {
		if v == col {
			return false
		}
	}
	if ok := pie[row+col]; ok {
		return false
	}
	if ok := na[row-col]; ok {
		return false
	}
	return true
}

func main() {
	ff := solveNQueens(4)
	fmt.Println(ff)
}
