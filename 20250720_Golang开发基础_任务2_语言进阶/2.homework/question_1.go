package main

import "fmt"

// 第一题 指针
// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func Question_1() {

	// vst := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 6, 5, 4, 3, 2, 1}

	// countMap := make(map[int]int)
	// for _, v := range vst {
	// 	countMap[v]++
	// }
	// fmt.Println(countMap)
	// for k, v := range countMap {
	// 	if v == 1 {
	// 		fmt.Println(k)
	// 	}
	// }

	var count int = 1
	changeNumber(&count)
	fmt.Println(count)

}

func changeNumber(num *int) { // 引用传递
	*num += 10
}
