package main

import "fmt"

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/

func main() {
	ch := make(chan int) // 创建一个整型通道

	go producer(ch) // 启动生产者协程
	consumer(ch)    // 启动消费者协程（在主协程中运行）
}

// 生成者协程：发送1到10的整数到通道
func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i // 发送数据到通道
	}
	close(ch) // 发送完成后关闭通道
}

// 消费者协程：从通道接收数据并打印
func consumer(ch chan int) {
	for num := range ch { // 通过 range 自动从通道接收，直到通道关闭
		fmt.Println("消费:", num)
	}
}
