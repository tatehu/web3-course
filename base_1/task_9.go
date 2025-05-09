package main

import "fmt"

/*
430. 扁平化多级双向链表：多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。给定位于列表第一级的头节点，
请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
可以定义一个结构体来表示链表节点，包含 val、prev、next 和 child 指针，然后使用递归的方法来扁平化链表，先处理当前节点的子链表，再将子链表插入到当前节点和下一个节点之间。
*/

func main() {
	head := buildTestList()
	flattened := flatten(head)
	print(flattened)
}

// 打印链表
func print(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

// 构建测试链表
func buildTestList() *Node {
	// 构造如下结构：
	// 1 - 2 - 3 - 4 - 5 - 6
	//         |
	//         7 - 8 - 9 - 10
	//             |
	//             11 - 12
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 9}
	n10 := &Node{Val: 10}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}

	n1.Next = n2
	n2.Prev = n1
	n2.Next = n3
	n3.Prev = n2
	n3.Next = n4
	n4.Prev = n3
	n4.Next = n5
	n5.Prev = n4
	n5.Next = n6
	n6.Prev = n5

	n3.Child = n7
	n7.Next = n8
	n8.Prev = n7
	n8.Next = n9
	n9.Prev = n8
	n9.Next = n10
	n10.Prev = n9

	n8.Child = n11
	n11.Next = n12
	n12.Prev = n11

	return n1
}

// Node 定义链表节点结构
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// flatten 扁平化多级双向链表
func flatten(head *Node) *Node {
	if head == nil {
		return nil
	}
	flattenDFS(head)
	return head
}

// flattenDFS 递归展开子链表，返回尾节点
func flattenDFS(node *Node) *Node {
	curr := node
	var last *Node

	for curr != nil {
		next := curr.Next
		// 如果当前节点有子链表
		if curr.Child != nil {
			childHead := curr.Child
			childTail := flattenDFS(childHead)

			// 插入子链表
			curr.Next = childHead
			childHead.Prev = curr
			curr.Child = nil

			if next != nil {
				childTail.Next = next
				next.Prev = childTail
			}
			last = childTail
			curr = next
		} else {
			last = curr
			curr = curr.Next
		}
	}
	return last
}
