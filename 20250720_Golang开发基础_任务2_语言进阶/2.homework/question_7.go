package main

import (
	"fmt"
	"time"
)

// 第七题 Channel
// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 终端运行    go run main.go Question_7.go

// 只接收channel的函数
func receive(ch <-chan int) {
	for i := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("接收到: %d\n", i)
	}
}

// 发送channel的函数
func send(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
func Question_7() {

	//创建一个channel
	ch := make(chan int, 10)

	//启动两个协程
	//接受协程序
	go receive(ch)
	//发送协程
	go send(ch)
	// time.Sleep(time.Second * 1)

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("通道已关闭")
				return
			}
			fmt.Printf("主携程接收到了%d\n:", v)

			return
		default:
			fmt.Println("没有数据,等待中....")
			time.Sleep(time.Second * 2)
		}

	}

}
