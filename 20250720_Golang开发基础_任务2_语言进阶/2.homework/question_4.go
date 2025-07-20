package main

import (
	"fmt"
	"time"
)

// 第四题 Goroutine
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，
// 同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

// task 函数返回 任务执行结果
type Task func() string

// taskResult 用于保存任务执行结果
type taskResult struct {
	name     string
	result   string
	duration time.Duration
}

func executeTask(taskName string, task Task, resultChan chan<- taskResult) {
	start := time.Now()
	result := task()
	duration := time.Since(start)
	resultChan <- taskResult{taskName, result, duration}

}

func TaskScheduler(tasks map[string]Task) {
	resultChan := make(chan taskResult, len(tasks))

	//启动 goroutine 执行每个任务
	for taskName, task := range tasks {
		go executeTask(taskName, task, resultChan)
	}

	//收集所有任务的执行结果
	for range tasks {
		res := <-resultChan
		fmt.Println(res.name, "执行完毕，耗时：", res.duration, "结果：", res.result)
	}

}

func Question_4() {

	tasks := map[string]Task{
		"任务A": func() string {
			time.Sleep(2 * time.Second)
			return "任务A执行完毕"
		},
		"任务B": func() string {
			time.Sleep(500 * time.Microsecond)
			return "任务B执行完毕"
		},
		"任务C": func() string {
			time.Sleep(3 * time.Second)
			return "任务C执行完毕"
		},
		"任务D": func() string {
			time.Sleep(900 * time.Microsecond)
			return "任务D执行完毕"
		},
	}

	fmt.Println("任务调度开始.........")
	start := time.Now()

	//启动调度器
	TaskScheduler(tasks)

	end := time.Since(start)
	fmt.Println("任务调度结束，耗时：", end)

}
