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

func isSubsequence(s string, t string) bool {
	len1 := len(s)
	len2 := len(t)

	if len1 > len2 {
		return false
	}
	ind1 := 0
	ind2 := 0

	for ind1 < len1 && ind2 < len2 {
		if s[ind1] == t[ind2] {
			ind1++
		}
		ind2++
	}

	return ind1 == len1
}

func twoSum(numbers []int, target int) []int {

	var res []int

	left := 0
	right := len(numbers) - 1

	for left != right {
		if numbers[left]+numbers[right] == target {
			res = append(res, left+1, right+1)

			break
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right++
		}
	}

	return res

}
