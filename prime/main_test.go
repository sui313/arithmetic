package main

import (
	"testing"
	"time"
)

func Test_GetPrimes(t *testing.T) {
	GetPrimes(10000)
}

func BenchmarkGetPrimes(b *testing.B) {
	b.StopTimer()
	time.Sleep(time.Millisecond * 500) // 模拟某个耗时但与被测程序关系不大的操作。
	max := 1000
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		GetPrimes(max)
	}
}
