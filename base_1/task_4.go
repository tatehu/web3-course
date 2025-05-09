package main

/*
46. 全排列：给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
可以使用回溯算法，定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列，使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
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
