package main

import (
	"fmt"
	"time"
)

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

// 发送channel的函数
func producer(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("生产者生产了一个数据：", i)
		time.Sleep(10 * time.Millisecond)
	}
	close(ch)
}

// 只接收channel的函数
func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Printf("消费者接收到了一个数据: %d\n", i)
		time.Sleep(20 * time.Millisecond)

	}
}

func Question_8() {

	//创建一个channel ,缓冲大小为 10
	ch := make(chan int, 10)

	//启动两个协程
	// 启动消费者协程
	go consumer(ch)

	//发送协程
	//启动生产者协程
	go producer(ch)
	// time.Sleep(time.Second * 1)

	time.Sleep(time.Second * 5)

	fmt.Println("主程序结束.........")

}
