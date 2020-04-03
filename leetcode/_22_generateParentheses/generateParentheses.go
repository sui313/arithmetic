/**
* @Author: sui_liut@163.com
* @Date: 2020/3/29 15:05
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var result []string
	_generate(0, 0, n, &result, "")
	return result
}

func _generate(left, rigth, n int, result *[]string, s string) {
	fmt.Println(s)
	if left == n && rigth == n {
		*result = append(*result, s)
		return
	}
	if left < n {
		_generate(left+1, rigth, n, result, s+"(")
	}
	if rigth < left {
		_generate(left, rigth+1, n, result, s+")")
	}
}
