package main

import "fmt"

/*
21. 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
可以定义一个函数，接收两个链表的头节点作为参数，在函数内部使用双指针法，通过比较两个链表节点的值，将较小值的节点添加到新链表中，直到其中一个链表为空，然后将另一个链表剩余的节点添加到新链表中。
*/

// 测试代码
func main() {
	// 创建测试链表1: 1 -> 3 -> 5
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	// 创建测试链表2: 2 -> 4 -> 6
	l2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}

	// 合并链表
	merged := mergeTwoLists(l1, l2)

	// 输出合并结果
	printList(merged)
}

// 定义链表节点结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个升序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个虚拟头节点，便于操作
	dummy := &ListNode{}
	current := dummy

	// 双指针法：遍历两个链表
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// 将剩余部分直接连接到新链表后面
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	// 返回合并后链表的头节点（跳过 dummy）
	return dummy.Next
}

// 打印链表辅助函数
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
