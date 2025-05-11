package main

import "fmt"

/*
729. 我的日程安排表 I：实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，
则可以存储这个新的日程安排。当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。日程可以用一对整数 start 和 end 表示，
这里的时间是半开区间，即 [start, end) ，实数 x 的范围为 start <= x < end 。实现 MyCalendar 类：MyCalendar() 初始化日历对象。boolean book(int start, int end)
如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true ，否则，返回 false 并且不要将该日程安排添加到日历中。
可以定义一个结构体来表示日程安排，包含 start 和 end 字段，然后使用一个切片来存储所有的日程安排，在 book 方法中，遍历切片中的日程安排，判断是否与要添加的日程安排有重叠。
*/

func main() {
	cal := Constructor()

	// 添加几个日程，观察是否成功
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(15, 25))
	fmt.Println(cal.Book(20, 30))
	fmt.Println(cal.Book(5, 10))
	fmt.Println(cal.Book(10, 15))
	fmt.Println(cal.Book(5, 15))
}

// 定义一个结构体表示日程安排
type Interval struct {
	start, end int
}

// 定义 MyCalendar 结构体
type MyCalendar struct {
	events []Interval
}

// 初始化 MyCalendar
func Constructor() MyCalendar {
	return MyCalendar{
		events: []Interval{},
	}
}

// 实现 book 方法
func (c *MyCalendar) Book(start int, end int) bool {
	for _, event := range c.events {
		// 如果两个时间区间有重叠，返回 false
		if max1(start, event.start) < min(end, event.end) {
			return false
		}
	}
	// 没有重叠，添加到日历中
	c.events = append(c.events, Interval{start, end})
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
