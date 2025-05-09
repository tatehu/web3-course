package main

import "fmt"

/*
69. x 的平方根：实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
可以使用二分查找法来解决，定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，直到找到满足条件的平方根或者确定不存在精确的平方根。
*/

func main() {
	fmt.Println(mySqrt(4)) // 输出：2
	fmt.Println(mySqrt(8)) // 输出：2（因为 sqrt(8) ≈ 2.828，只保留整数部分）
	fmt.Println(mySqrt(0)) // 输出：0
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	for left <= right {
		mid := left + (right-left)/2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
