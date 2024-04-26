package main

import (
	"errors"
	"fmt"
	"sync"
)

/*
*
go 语言函数笔记
go 语言函数支持普通函数、匿名函数、闭包

go 语言中的传递是值传递
*/
func main() {

	// go中函数是一等公民
	// 函数本身可以当做变量
	// 匿名函数、闭包
	// 函数可以满足接口
	a := 5
	b := 7
	fmt.Println(add(a, b))
	fmt.Println(a, b)

	// 函数本身可以当做变量
	funcVar := add1
	sum := funcVar(a, b, 4)
	fmt.Println(sum)

	//nextFund := autoIncres()
	//for i := 0; i < 4; i++ {
	//	fmt.Println(nextFund())
	//}
	defer_fmt_learn()

	_, err := error_learn()

	if err != nil {
		fmt.Println(err)
	}

	recover_learn()
}

// func add(a,b int) int{}
func add(a int, b int) int {
	a = a + 1
	return a + b
}

// 动态参数列表
func add1(item ...int) (sum int) {

	for _, value := range item {
		sum += value
	}
	return sum
}

// 函数作为返回类型,闭包
func autoIncres() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}

//函数defer的应用场景
/**
数据库连接、打开文件、开始锁，无论如何 最后都要记得去关闭数据库、关闭文件、解锁
*/
func defer_learn() {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock() //defer 后面的代码是会放在函数return之前执行
	/**
	打开文件

	逻辑
	*/
}

func defer_fmt_learn() {

	//执行顺序由下至上执行，类似与栈，先入栈后执行
	defer fmt.Println("1")
	defer fmt.Println("2")

	fmt.Println("main")
	return
}

// error go语言函数需要一个返回值去告诉调用者是否成功，go设计者必须要处理这个err
func error_learn() (int, error) {

	return 1, errors.New("this is a error")
}

// panic 这个函数，会导致程序退出
func panic_learn() {
	// 开发中不要随便用
	// 当服务想要启动，其依赖的服务需要准备好，如果没有启动，需要调用panic
	// 服务一旦启动，执行到panic，服务就会挂掉
	panic("this is panic")

	fmt.Println("hello")
}

// recover 这个函数能够捕获到panic

func recover_learn() {

	// recover 只有在defer调用的函数才会生效
	// recover 处理异常后，并不会恢复到panic语句中
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover执行", r)
		}
	}()
	panic("this is panic")
	fmt.Println("hello")
}
