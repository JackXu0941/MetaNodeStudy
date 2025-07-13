package main

import (
	"fmt"
)

// 第三题
// 字符串
// 有效的括号
// 考察：字符串处理、栈的使用
// 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

func Question_3() bool {

	//回文数例子
	// vst := "{[()]}"     // true
	// vst := "{[()]}(){}[]((([[[]]])))"  // true
	vst := "{[()]}(){}[]((([[[]]])))[[]" //false

	n := len(vst)

	if n%2 == 1 {
		fmt.Println("不是有效括号字符串!")
		return false
	}

	pire := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}

	for i := 0; i < n; i++ {
		if pire[vst[i]] > 0 {
			if len(stack) == 0 || pire[vst[i]] != stack[len(stack)-1] {
				fmt.Println("不是有效括号字符串!")
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, vst[i])
		}

	}

	if len(stack) == 0 {
		fmt.Println("有效括号字符串!")
		return true
	} else {
		fmt.Println("不是有效括号字符串!")
		return false
	}

}
