package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 第十题 锁机制
// 	题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，
// 	最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

func Question_10() {

	var counter int32 = 0

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}

	time.Sleep(time.Second * 5)

	fmt.Println("计数结果：", counter)

	fmt.Println("主程序结束.........")

}
