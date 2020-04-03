package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quick(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	i, j := 0, n-1
	t := i
	for i != j {
		for arr[j] > arr[t] && j > i {
			j--
		}
		for arr[i] <= arr[t] && i < j {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[t],arr[i] = arr[i],arr[t]
	quick(arr[:i])
	quick(arr[i+1:])
	return arr
}

func main()  {
	r:=rand.NewSource(time.Now().UnixNano())

	n:=10
	Arr :=make([]int,n,n)
	for i:=0;i<n;i++{
		Arr[i]=rand.New(r).Intn(100)
	}

	b:=[]int{5,1,1,2,0,0}

	fmt.Println(Arr)
	fmt.Println(quick(Arr))
	fmt.Println(quick(b))
}