package main

import "fmt"

/**
定义结构体
*/

func main() {

	// type 关键字
	// 1、定义结构体 2、定义接口 3、定义类型别名 4、类型定义

	// 定义类型别名
	type MyInt = int
	var i MyInt = 3
	fmt.Println(i)

	//类型定义
	type MyInt1 int
	var j MyInt1 = 3
	fmt.Printf("%T\r\n", j)

}
