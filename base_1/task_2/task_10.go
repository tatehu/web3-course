package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func main() {
	wg01.Add(10)
	for i := 0; i < 10; i++ {
		go increment01()
	}
	wg01.Wait()
	fmt.Println("最终计数器的值是：", counter01)
}

var (
	counter01 int64 // 使用 int64 类型，以支持 atomic 操作
	wg01      sync.WaitGroup
)

func increment01() {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counter01, 1) // 原子递增
	}
	wg01.Done()
}
