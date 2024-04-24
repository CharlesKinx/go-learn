package main

import "fmt"

func main() {

	// 变量的作用域

	// 匿名变量，定义了，但可以不使用
	var _ int = 10

	// iota 特殊常量，可以认为是一个可以被编译器修改的常量

	const (
		ERR1 = iota + 2
		ERR2
		ERR3
		ERR4
		ERR5 = "hello"
	)

	fmt.Println(ERR1, ERR2, ERR3, ERR4)
}
