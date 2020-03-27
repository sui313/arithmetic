package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	NUM := 20
	var arr = make([]int, NUM)
	for i := 0; i < NUM; i++ {
		arr[i] = r.Intn(100)
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] < arr[j] {
			return true
		}
		return false
	})
	arr2 := []int{2, 3, 6, 7, 15, 17, 19, 21, 31, 36, 37, 45, 48, 52, 61, 76, 76, 82, 87, 91}
	fmt.Println(arr2)
	fmt.Println(find(3, arr2))
}

/*
时间复杂度计算
假设元素有n个,查找次数为n/2,n/2^2,n/2^3 ...n/2^k,
到n/2^k=1时之剩下一个元素为止，k=log2(n)
*/
func find(n int, arr []int) []int {
	hig := len(arr) - 1
	low := 0
	mid := (hig + low) / 2
	if n < arr[low] || n > arr[hig] {
		fmt.Println("不存在", n)
		return nil
	}
	fmt.Println(n, arr[low], arr[hig])
	result := make([]int, 0)
	for low <= hig { //1个元素的情况
		if arr[low] == n {
			for low <= hig {
				if arr[low] != n {
					break
				}
				result = append(result, low)
				low++
			}
			return result
		}
		if arr[hig] == n {
			for hig >= low {
				if arr[hig] != n {
					break
				}
				result = append(result, hig)
				hig--
			}
			return result
		}

		if arr[mid] == n {
			t := mid - 1
			result = append(result, mid)
			mid = mid + 1
			for arr[mid] == n {
				result = append(result, mid)
				mid++
			}
			mid = t
			for arr[mid] == n {
				result = append(result, mid)
				mid--
			}
			return result
		}
		if n > arr[low] && n < arr[mid] {
			hig = mid - 1
			mid = (low + hig) / 2
		} else {
			low = mid + 1
			mid = (low + hig) / 2
		}
	}
	return result
}
