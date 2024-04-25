package main

import (
	"fmt"
	"strconv"
	"strings"
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
func isPalindrome(s string) bool {

	var str = strings.ToLower(getStr(s))
	var right = len(str)
	var left = 0

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func getStr(s string) string {
	var ans = ""
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			ans += string(s[i])
		}
	}
	return ans
}
