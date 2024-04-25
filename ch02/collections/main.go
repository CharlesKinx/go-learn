package main

import (
	"container/list"
	"fmt"
)

func main() {

	// go 数组
	// array_test()

	// go切片
	//slice_learn()
	// testSlice()

	//map_learn()

	list_learn()
}

/*
*
go 语言list学习
*/
func list_learn() {
	var list1 list.List

	list1.PushBack(1)
	list1.PushBack(2)
	list1.PushBack(3)
	list1.PushFront(4)
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Printf("%d", i.Value)
	}
	fmt.Println()
	// 插入数据

	ele := list1.Front()

	for ; ele != nil; ele = ele.Next() {
		if ele.Value == 2 {
			break
		}
	}

	list1.InsertAfter(9, ele)
	list1.InsertBefore(0, ele)

	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Printf("%d", i.Value)

	}
}

/**
go map学习

key value 无序数据结构，主要是查询方面
*/

func map_learn() {
	var hashmap = map[string]string{}

	hashmap["hello"] = "world"
	hashmap["he"] = "wo"
	hashmap["go"] = "gin"

	for k, v := range hashmap {
		fmt.Println(k, v)
	}
	// 判断元素是否存在
	v, ok := hashmap["h"]

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not exists")
	}

	//删除一个元素
	delete(hashmap, "go")
	fmt.Println(hashmap)

	//map不是线程安全的
}

/**
go切片学习
与数组初始化相同，不同的点不指定长度
var name []int
*/

func slice_learn() {
	var courses []int

	//像切片中添加数据，返回值必须要有切片
	courses = append(courses, 1)
	courses = append(courses, 2)
	courses = append(courses, 3, 4, 5, 6)

	fmt.Println(courses)

	// 切片的初始化
	// 1、从数组直接创建
	// 2、使用string{}
	// 3、使用make

	// 1、从数组直接创建
	var courses1 [3]string // 3个元素数组
	courses1[0] = "hello"
	courses1[1] = "world"
	courses1[2] = "从数组直接创建"
	courses2 := courses1[0:2]
	fmt.Println(courses2)

	// 2、使用string{}
	courses3 := []string{"hello", "world", "使用string"}
	courses3 = append(courses3, courses2...)
	fmt.Println(courses3)

	// 3、使用make,现指定切片大小
	courses4 := make([]string, 3)
	courses4[0] = "hello"
	courses4[1] = "world"
	//courses4[2] = "使用make"
	fmt.Println(courses4)

	// 访问切片的元素[start:end], 如果没有end，说明从start到结束
	fmt.Println("访问切片元素")
	fmt.Println(courses3[1:3])

	// 删除slice中的元素
	courses5 := []string{"hello", "world", "使用string", "wsx", "qaz"}
	mySlice := append(courses5[0:2], courses5[3:]...)
	fmt.Println(mySlice)

}

func testSlice() {
	courses := []string{"hello", "wsx", "qaz"}
	fmt.Println(courses)

	//courses1 := courses[0:3]
	//fmt.Println(courses1)
	//courses1[1] = "qqq"
	//fmt.Println(courses1)

	printSlice(courses)
	fmt.Println(courses)
}

func printSlice(data []string) {
	data[0] = "qaz"
	num := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	num1 := num[4:7]
	fmt.Println(num)
	fmt.Println(num1)
	fmt.Printf("SliceA len: %v, cap: %v\n", len(num1), cap(num1))
	num1 = append(num1, 4)
	fmt.Println(num)
	fmt.Println(num1)
	fmt.Printf("SliceA len: %v, cap: %v\n", len(num1), cap(num1))
	data = append(data, "q")
	fmt.Println(data)

}

/**
go数组的定义
var name [count]int
*/

func array_test() {

	var courses [3]string // 3个元素数组
	//var courses1 [4]string

	courses[0] = "hello"
	courses[1] = "world"
	courses[2] = "lv"

	fmt.Println(courses, len(courses))
	//fmt.Printf("%T\r\n", courses)
	//fmt.Printf("%T\r\n", courses1)

	for index, value := range courses {
		fmt.Println(index, value)
	}

	// 数组的初始化
	//var arr [4]string = [4]string{"a", "b", "c", "d"}
	arr := [4]string{"a", "b", "c", "d"}
	fmt.Println(arr, len(arr))
	arr1 := [...]string{"a", "b", "c", "d"}
	fmt.Println(arr1, len(arr1))

	//数组可以直接比较
	if arr == arr1 {
		fmt.Println("equal")
	}

	// 多维数组
	var arr2 [2][3]string
	arr2[0] = [3]string{"a", "b", "c"}
	arr2[1][0] = "q"
	arr2[1][1] = "a"
	arr2[1][2] = "z"

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

}
