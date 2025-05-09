package main

import "fmt"

/*
46. 全排列：给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
可以使用回溯算法，定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列，使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
*/

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println("全排列结果：", result)
}

func permute(nums []int) [][]int {
	var res [][]int
	backtrack(nums, 0, &res)
	return res
}

func backtrack(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		// 需要复制一份当前排列，否则后续的修改会影响结果
		perm := make([]int, len(nums))
		copy(perm, nums)
		*res = append(*res, perm)
		return
	}

	for i := start; i < len(nums); i++ {
		// 交换当前位置和当前遍历到的位置
		nums[start], nums[i] = nums[i], nums[start]
		// 递归处理后续的元素
		backtrack(nums, start+1, res)
		// 撤销交换，回溯
		nums[start], nums[i] = nums[i], nums[start]
	}
}
