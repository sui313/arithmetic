/*
10000以内的素数
*/
package main

import (
	"fmt"
	"runtime"
)

//import "sync/atomic"
import "time"
import (
	"context"
)

var count int32

func main() {
	if false {
		ctx := context.WithValue(context.Background(), "test111", "123")
		ctx, cancel := context.WithCancel(ctx)
		//ctx, cancel :context.WithTimeout(context.Background(), 2*time.Second)
		for i := 0; i < 10; i++ {
			go tt2(ctx, i)
		}

		time.Sleep(500 * time.Millisecond)
		cancel()
		time.Sleep(10 * time.Second)
	}
	fmt.Println("cpu:", runtime.NumCPU())
	GetPrimes(100)
}

func tt2(ctx context.Context, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done:", i, ctx.Err())
			return
		default:
			fmt.Println("1----", i, ctx.Value("test111"))
			time.Sleep(400 * time.Millisecond)
			fmt.Println("2----", i)
		}
	}
}

/*
先找小的素数，然后用后面的数除以找出的小的素数.
如果整除就不是素数。
*/
func GetPrimes(max int) {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < max; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		//atomic.AddInt32(&count, 1)
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		//fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			//time.Sleep(500 * time.Millisecond)
			//fmt.Println(num, "-----", prime)
			if num%prime != 0 { //被整除的就不是素数过滤掉，没被整除的继续传递给下一个素数去整除
				out <- num
			}
		}
		//fmt.Println("-----", prime)
		close(out)
	}()
}
