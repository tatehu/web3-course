package main

import (
	"fmt"
	"sort"
)

/*
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/

func main() {
	nums := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {2, 3}},
	}

	for i, intervals := range nums {
		fmt.Printf("测试案例 %d: 输入: %v\n", i+1, intervals)
		result := merge(intervals)
		fmt.Printf("输出: %v\n\n", result)
	}
}

// 定义区间类型为二维整数数组
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照每个区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 用一个切片存储合并后的区间
	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		// 判断是否重叠：当前区间的 start <= 上一个合并区间的 end
		if current[0] <= last[1] {
			// 合并：更新 end 为两者中较大的
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 不重叠，直接加入结果集
			merged = append(merged, current)
		}
	}

	return merged
}
