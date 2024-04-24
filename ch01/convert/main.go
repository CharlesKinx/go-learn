package main

import (
	"fmt"
	"strconv"
)

func main() {

	var istr = "123"
	var myInt, err = strconv.Atoi(istr)

	if err != nil {
		fmt.Println("转化失败")
	}

	fmt.Println(myInt)

	var myI = 123

	var mystr = strconv.Itoa(myI)
	fmt.Println(mystr)

	// 字符串转化为float
	var s = "3.1254"
	var myf, e = strconv.ParseFloat(s, 64)
	if e != nil {
		fmt.Println("转化失败")
	}
	fmt.Println(myf)

	// 字符串转int

	var s1 = "123"
	var myi, e1 = strconv.ParseInt(s1, 10, 64)
	if e1 != nil {
		fmt.Println("转化失败")
	}
	fmt.Println(myi)
}
