package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
算法题： 在1个10G大小的文件中，存储的都是int型的数据，如何在内存使用小于8M的情况下进行排序
*/
var threshold = 2000000

func main() {
	Init()
	cutFile("test.txt")
	return
	arr := &[]int{1, 2}

	fmt.Println(arr)

	ap(arr, []int{1, 5, 1, 1, 1, 1, 1, 233, 1, 1, 1, 1})
	sort.Ints(*arr)
	fmt.Println(arr)
}

func ap(arr *[]int, a []int) {
	*arr = append(*arr, a...)
}

func Init() {
	start := time.Now()
	defer func() {
		fmt.Println("创建数据花费:", time.Now().Sub(start))
	}()

	news := rand.NewSource(time.Now().UnixNano())
	r := rand.New(news)
	var arr = make([]int, threshold)
	f, _ := os.Create("test.txt")
	defer f.Close()
	for i := 0; i < 5; i++ {
		for j := 0; j < threshold; j++ {
			arr[j] = int(r.Int31n(1000))
		}
		//sort.Ints(arr)
		writeFile(&arr, f)
	}
}

type ReadFile struct {
	F     *os.File
	Index int
	Arr   []int
}

var FM = make(map[string]*ReadFile)

func cutFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件失败:", err.Error())
		return err
	}

	nRead := bufio.NewReader(f)
	var arr *[]int
	arr = &[]int{}
	var id int
	var fw *os.File
	for {
		if fw == nil {
			fn := fmt.Sprintf("file_%d.txt", id)
			fw, err = os.Create(fn)
			if err != nil {
				fmt.Println("创建文件失败:", err.Error())
				return err
			}
			id++
			FM[fn] = nil
		}
		line, err := nRead.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("EOF:文件名:%s,数组长度:%d,数组地址:%p\n", fw.Name(), len(*arr), arr)
				numSort(line, arr, fw)
				fw.Close()
				break
			} else {
				fw.Close()
				fmt.Println("读取文件出错:", err.Error())
				break
			}
		}

		if numSort(line, arr, fw) {

			fmt.Printf("文件名:%s,数组长度:%d,数组地址:%p\n", fw.Name(), len(*arr), arr)
			fw.Close()
			fw = nil
			arr = &[]int{}
			fmt.Println("新数组")
		}
	}
	return nil
}

func numSort(line string, arr *[]int, f *os.File) bool {

	tmp := strings.Split(line, ",")
	for i, _ := range tmp {
		if tmp[i] == "" {
			continue
		}
		num, _ := strconv.Atoi(tmp[i])
		if num == 0 {
			fmt.Printf("num=%v,tmp[%d]=%v,file=%s\n", num, i, tmp[i], f.Name())
		}
		*arr = append(*arr, num)
	}
	fmt.Println((*arr)[0])
	if len(*arr) >= threshold {
		sort.Ints(*arr)
		writeFile(arr, f)
		return true
	}
	return false
}

func writeFile(arr *[]int, f *os.File) {

	var s string

	for k, _ := range *arr {
		s = strconv.Itoa((*arr)[k])
		f.WriteString(s)
		if k+1 != len(*arr) {
			f.WriteString(",")
		}
		if k%100 == 0 && k > 0 {
			f.WriteString("\n")
		}
	}
}
