package main

import (
	"fmt"
	"sync"
	"time"
)

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

func main() {
	tasks := []Task{
		{
			Name: "task1",
			Func: func() {
				time.Sleep(2 * time.Second)
				fmt.Println("任务 1 完成")
			},
		},
		{
			Name: "task2",
			Func: func() {
				time.Sleep(1 * time.Second)
				fmt.Println("任务 2 完成")
			},
		},
		{
			Name: "task3",
			Func: func() {
				fmt.Println("任务 3 完成（无延迟）")
			},
		},
	}

	ExecuteTasks(tasks)
}

// 定义一个任务类型
type Task struct {
	Name string
	Func func()
}

// 并发执行任务，并统计执行时间
func ExecuteTasks(tasks []Task) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make(map[string]time.Duration)

	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()

			start := time.Now()
			t.Func()
			duration := time.Since(start)

			// 保护共享资源
			mu.Lock()
			results[t.Name] = duration
			mu.Unlock()
		}(task)
	}

	wg.Wait()

	fmt.Println("任务执行时间统计：")
	for name, duration := range results {
		fmt.Printf("任务 %s 执行时间: %v\n", name, duration)
	}
}
