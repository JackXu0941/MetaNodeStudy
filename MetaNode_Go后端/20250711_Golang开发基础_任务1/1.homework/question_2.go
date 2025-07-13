package main

import (
	"fmt"
)

// 第二题
// 回文数
// 考察：数字操作、条件判断
// 题目：判断一个整数是否是回文数

func Question_2() {

	//回文数例子
	vst := []string{"a", "b", "c", "b", "a"}
	//非回文数例子
	// vst := []string{"a", "b", "c", "b", "a", "a1"}

	flag := true
	count := int(len(vst) / 2)
	length := len(vst) - 1

	fmt.Println(count)
	for i := 0; i <= count; i++ {
		fmt.Println(i, count)

		if vst[i] != vst[length-i] {
			flag = false
		}
	}

	if flag {
		fmt.Println("回文数")
	} else {
		fmt.Println("非回文数")
	}

}
