package main

import "fmt"

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

func main() {
	// 测试用例
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", numbers)

	// 调用函数修改切片
	multiplyByTwo(&numbers)

	// 打印结果验证
	fmt.Println("乘以2后的切片:", numbers)
}

func multiplyByTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}
