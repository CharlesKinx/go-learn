package main

import "fmt"

// 全局变量,可以不使用
var name = "global"

func main() {

	// go是静态语言
	// 1、变量必须先定义后使用，变量必须有类型，类型定下来后不能改变

	// go语言中，变量定义了必须要使用
	var age int = 25
	name := "hello world"

	fmt.Println(name)
	fmt.Println(age)

	//常量，定义的时候就指定的值，不能修改
	// 命名需要大写加下划线进行命名

	//显示定义常量
	const PI float32 = 3.1415
	fmt.Println(PI)

	// 隐式定义不需要说明数据类型
	const (
		UNKNOW = 1
		MALE   = 2
	)

	/**
	常量类型只可以定义基本类型和字符串
	不使用的常量，不会报错
	显示指定类型的时候，必须确保值的类型是一样的
	*/
	const (
		x int = 16
		y
		s = "123"
		z
	)

	fmt.Println(y)
}
