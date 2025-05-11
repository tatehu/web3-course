package main

import (
	"fmt"
	"sync"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printOdd(&wg)
	go printEven(&wg)

	wg.Wait()
}

func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("Odd: %d\n", i)
	}
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("Even: %d\n", i)
	}
}
