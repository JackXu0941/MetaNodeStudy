package main

import "fmt"

// 第五题
// 难度：简单
// 考察：数组操作、进位处理
// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

func Question_5() []int {

	//实例1 返回[1,2,4]
	// digits := []int{1, 2, 9, 9, 9}
	//实例1 返回[1,3,0,0,0]
	digits := []int{9, 9, 9, 9, 9}

	result := plusOne(digits)

	fmt.Println(result)
	return result

}

func plusOne(digits []int) []int {

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++

			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}

	}

	digits = make([]int, n+1)
	digits[0] = 1
	return digits

}
