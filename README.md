# go-learn



### *四月*

|             Mon.             |             Tues.             | Wed. | Thur. |             Fri.             |             Sat.             |             Sun.             |
| :--------------------------: | :---------------------------: | :--: | :---: | :--------------------------: | :--------------------------: | :--------------------------: |
|              1               |               2               |  3   |   4   |              5               |              6               |              7               |
|              8               |               9               |  10  |  11   |              12              |              13              |              14              |
|              15              |              16               |  17  |  18   |              19              |              20              |              21              |
|              22              |              23               |  24  |  25   | 26<br>([D1](#2024426-Day1 )) | 27<br>([D2](#2024427-Day2 )) | 28<br>([D3](#2024428-Day3 )) |
| 29<br>([D4](#2024429-Day4 )) | 29<br/>([D5](#2024430-Day5 )) |      |       |                              |                              |                              |

### *五月*

| Mon. | Tues. |             Wed.             | Thur. | Fri. | Sat. | Sun. |
| :--: | :---: | :--------------------------: | :---: | :--: | :--: | :--: |
|      |       | 1<br/>([D6](#2024501-Day6 )) |   2   |  3   |  4   |  5   |
|  6   |   7   |              8               |   9   |  10  |  11  |  12  |
|  13  |  14   |              15              |  16   |  17  |  18  |  19  |
|  20  |  21   |              22              |  23   |  24  |  25  |  26  |
|  27  |  28   |              29              |  30   |  31  |      |      |



##  2024/4/26 Day1 

今天学习时长1h

主要学习内容：

- 函数的基本使用
- 函数的动态参数
- 函数可以作为参数以及返回值
- defer、panic以及recover的使用

## 2024/4/27 Day2 

leetcode 刷题

- [盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&amp;envId=top-interview-150) 
- [三数之和](https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&amp;envId=top-interview-150)

今天学习内容：

- go语言的结构体

## 2024/4/28 Day3

leetcode 刷题

- [长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)

go语言的指针用法

## 2024/4/29 Day4

LeetCode刷题

[无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)

## 2024/4/30 Day5

今天摆烂...

## 2024/5/1 Day6

LeetCode刷题

[串联所有单词的子串](https://leetcode.cn/problems/substring-with-concatenation-of-all-words/)

go语言中没有用来比较两个hashmap是否相等，今天写了一个比较hashmap的函数

```go
func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if _, ok := b[k]; !ok || v != b[k] {
			return false
		}
	}
	return true
}
```



