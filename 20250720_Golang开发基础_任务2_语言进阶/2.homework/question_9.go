package main

import (
	"fmt"
	"sync"
	"time"
)

// 第九题 锁机制
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，
// 最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

// 线程安全计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 获取计数
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func Question_9() {

	counter := SafeCounter{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}()
	}

	time.Sleep(time.Second * 5)

	fmt.Println("计数结果：", counter.Value())

	fmt.Println("主程序结束.........")

}
