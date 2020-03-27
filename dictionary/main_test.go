package main

import (
	"fmt"
	"testing"
)

func TestFindkvTest(t * testing.T)  {
	var m1 =make(map[string]interface{})
	var m2 =make(map[string]interface{})

	m1["A"]=1
	m1["B"]=map[string]interface{}{
		"C":3,
		"D":4,
	}

	m2["B"]  = 1234

	m3:=findKV(m1,m2)
	fmt.Println(m3)
}