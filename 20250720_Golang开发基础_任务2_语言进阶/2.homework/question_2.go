package main

import "fmt"

// 第二题 指针
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func Question_2() {

	vst := []int{1, 2, 3, 4, 5}
	multiplyNumber(&vst)
	for i, v := range vst {
		fmt.Println(i, v)
	}

}

func multiplyNumber(num *[]int) { // 引用传递

	for i, _ := range *num {
		(*num)[i] *= 2
	}
}
