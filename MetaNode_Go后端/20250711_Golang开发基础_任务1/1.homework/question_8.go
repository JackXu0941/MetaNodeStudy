package main

import (
	"fmt"
)

// 第八题目
// 两数之和
// 考察：数组遍历、map使用
// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func Question_8() []int {

	//实例1 返回
	digits := []int{1, 2, 5, 6, 7, 8, 11, 15, 22}
	target := 10

	hashTable := map[int]int{}
	for i, x := range digits {

		if value, ok := hashTable[target-x]; ok {
			fmt.Println("目标值是:", target)
			fmt.Println("数组下标是:", value, i)
			fmt.Println("元素对应值:", digits[value], digits[i])
			return []int{value, i}
		}

		hashTable[x] = i
	}

	return nil
}
