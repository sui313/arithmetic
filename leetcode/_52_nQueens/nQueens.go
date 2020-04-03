/**
* @Author: sui_liut@163.com
* @Date: 2020/3/31 0:21
 */

package main

import "fmt"

func main() {
	x:=12345
	a:=23
	b:=-1113
	fmt.Println(a&b)
	fmt.Printf("%b\n",x)
	fmt.Printf("%b\n",-12345)
	fmt.Println(x&-x)
	fmt.Println(x&(x-1))
	// -1  1000 0001
}
