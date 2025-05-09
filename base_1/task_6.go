package main

/*
69. x 的平方根：实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
可以使用二分查找法来解决，定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，直到找到满足条件的平方根或者确定不存在精确的平方根。
*/

func main() {
	nums := []int{2, 2, 1}
	println(SingleNumber(nums))
}

func SingleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
