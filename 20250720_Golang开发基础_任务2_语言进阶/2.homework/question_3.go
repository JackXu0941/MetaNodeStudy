package main

import (
	"fmt"
	"time"
)

// 第三题 Goroutine
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，
// 另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func Question_3() {

	// // vst := []int{1, 2, 3, 4, 5}
	// // multiplyNumber(&vst)
	// // for i, v := range vst {
	// // 	fmt.Println(i, v)
	// // }

	go function1()
	go function2()
	time.Sleep(2000 * time.Millisecond)

	// go func() {
	// 	fmt.Println("run goroutine in closure")
	// }()
	// go func(string) {
	// }("gorouine: closure params")
	// go say("in goroutine: world")
	// say("hello")

}

func function1() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("奇数", 2*i-1)

	}
}

func function2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("偶数", 2*i)

	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
