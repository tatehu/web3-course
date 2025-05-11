package main

import (
	"fmt"
)

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

func main() {
	// 创建一个带缓冲的通道，缓冲大小可以调节性能
	bufferSize := 10
	ch := make(chan int, bufferSize)

	// 启动生产者和消费者协程
	go producer01(ch)
	consumer01(ch) // 主协程调用消费者函数，阻塞直到接收完成
}

// 生产者：向通道发送 100 个整数
func producer01(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch) // 发送完成后关闭通道
}

// 消费者：从通道接收并打印
func consumer01(ch chan int) {
	for val := range ch {
		fmt.Println("消费:", val)
	}
}
