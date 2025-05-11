package main

import "fmt"

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者
*/

func main() {
	emp := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
		},
		EmployeeID: "E12345",
	}

	emp.PrintInfo()
}

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Employee 结构体，组合了 Person
type Employee struct {
	Person     // 匿名字段，实现组合
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("EmployeeID: %s\n", e.EmployeeID)
}
