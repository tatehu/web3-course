package main

import (
	"fmt"
	"math"
)

/*
 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

func main() {
	// 创建 Rectangle 实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("Rectangle:")
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// 创建 Circle 实例
	circ := Circle{Radius: 4}
	fmt.Println("\nCircle:")
	fmt.Printf("Area: %.2f\n", circ.Area())
	fmt.Printf("Perimeter: %.2f\n", circ.Perimeter())

	// 使用 Shape 接口测试
	var s Shape

	s = rect
	fmt.Println("\nShape Rectangle:")
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())

	s = circ
	fmt.Println("Shape Circle:")
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// 定义 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义 Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现 Rectangle 的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现 Rectangle 的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义 Circle 结构体
type Circle struct {
	Radius float64
}

// 实现 Circle 的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现 Circle 的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
