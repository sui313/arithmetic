/*
Given a “flatten” dictionary object, whose keys are dot-separated. For example, { ‘A’: 1, ‘B.A’: 2, ‘B.B’: 3, ‘CC.D.E’: 4, ‘CC.D.F’: 5}.
Implement a function in any language to transform it to a “nested” dictionary object. In the above case, the nested version is like:
{
‘A’: 1,
	‘B’: {
		‘A’: 2,
		‘B’: 3,
	},
‘CC’: {
	‘D’: {
		‘E’: 4,
		‘F’: 5,
	}
	}
}
It’s guaranteed that no keys in dictionary are prefixes of other keys.
*/

package main

import (
	"fmt"
	"log"
	"strings"
	"encoding/json"
)

func main() {
	s := `{
    "A": 1,
    "B.A": 2,
    "B.B": 3,
    "CC.D.E": 4,
    "CC.D.F": 5,
    "DD.AA.BB.CC.EE.F1": 6,
    "DD.AA.BB.CC.EE.F2": 7,
    "DD.AA.BB.CC.EE.F3": 8,
	"DD.AA.CC": 9,
	"DD.AA.DD": 10
	}`
	nested(s)
	return
}

func nested(s string) {
	var m = make(map[string]interface{})

	err := json.Unmarshal([]byte(s), &m) //解析原始数据
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("===========1111:", m)
	var m2 = make(map[string]interface{})
	for key, val := range m {
		m3 := splitKey(key, val)
		fmt.Println("===========M3:",key,val, m3)
		for k, _ := range m3 {//遍历新map
			if _, ok := m2[k]; ok {//查找key是否存在m2中
				m2[k] = findKV(m2[k], m3[k]) //key已经在m2中，继续判断是否有嵌套的多级其它key
			} else {//不存在直接加入到m2字典中
				m2[k] = m3[k]
			}
		}
	}
	fmt.Println("===========", m)
	fmt.Println("===========", m2)
	js, _ := json.Marshal(m2)
	fmt.Println("===========", string(js))
}

//分割key
func splitKey(key string, v interface{}) map[string]interface{} {
	k := strings.Split(key, ".")
	l := len(k)
	if l == 0 {
		return nil
	}
	for i := l - 1; i > -1; i-- {
		v = kv(k[i], v)
	}
	return v.(map[string]interface{})
}

//返回key=>value字典
func kv(key string, v interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m[key] = v
	return m
}

//查找已存在的key,一层层的找
func findKV(m1, m2 interface{}) map[string]interface{} {
	t1, ok := m1.(map[string]interface{})
	if !ok {
		return nil
	}
	t2, ok := m2.(map[string]interface{})
	if !ok { //m2的值不符合规范 不把m2加入字典
		return t1
	}
	for k2, v2 := range t2 {
		if _, ok := t1[k2]; ok {
			t1[k2] = findKV(t1[k2], t2[k2])
		} else { //在当前层级下k2,v2不存在，直接把k2,v2加入到t1
			t1[k2] = v2
		}
	}
	return t1
}