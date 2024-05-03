# go-learn



#### *四月*

|             Mon.             | Tues. | Wed. | Thur. |             Fri.             |             Sat.             |             Sun.             |
| :--------------------------: | :---: | :--: | :---: | :--------------------------: | :--------------------------: | :--------------------------: |
|              1               |   2   |  3   |   4   |              5               |              6               |              7               |
|              8               |   9   |  10  |  11   |              12              |              13              |              14              |
|              15              |  16   |  17  |  18   |              19              |              20              |              21              |
|              22              |  23   |  24  |  25   | 26<br>([D1](#2024426-Day1 )) | 27<br>([D2](#2024427-Day2 )) | 28<br>([D3](#2024428-Day3 )) |
| 29<br>([D4](#2024429-Day4 )) |  30   |      |       |                              |                              |                              |



#### *五月*

| Mon. | Tues. |             Wed.             | Thur. |             Fri.             | Sat. | Sun. |
| :--: | :---: | :--------------------------: | :---: | :--------------------------: | :--: | :--: |
|      |       | 1<br/>([D5](#2024501-Day6 )) |   2   | 3<br/>([D6](#2024503-Day7 )) |  4   |  5   |
|  6   |   7   |              8               |   9   |              10              |  11  |  12  |
|  13  |  14   |              15              |  16   |              17              |  18  |  19  |
|  20  |  21   |              22              |  23   |              24              |  25  |  26  |
|  27  |  28   |              29              |  30   |              31              |      |      |

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

今天摆烂...没有学..

## 2024/5/01 Day6

突然发现....昨天写的没了...草！！！

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



go modules

```
go get "框架地址" 用于下载包，go get会修改go.mod文件

go mod用法

download    download modules to local cache
edit        edit go.mod from tools or scripts
graph       print module requirement graph
init        initialize new module in current directory
tidy        add missing and remove unused modules
vendor      make vendored copy of dependencies
verify      verify dependencies have expected content
why         explain why packages or modules are needed

go get -u 用来升级最新版本
```

### 代码规范

#### 1、代码规范

##### 1.1 命名规范

- 包名

  1、尽量和目录保持一致

  2、尽量采取有意义的包名

  3、不要和标注库名冲突

  4、包名采用全部小写

- 文件名

  user_name.go 小写加下划线

- 变量名

  尽量驼峰命名

- 结构体命名

  User

- 接口命名，和结构体命名差不多

- 常量命名，全部大写，如果有多个单词，使用下划线QUESTION_ANSWER

#### 2、注释规范

// 单行注释

/**

多行注释

/

#### 3、import规范

1、go自带的包

2、第三方的包

3、内部的包

#### 单元测试

go test 命令是一个按照一定约定和组织的测试代码驱动程序，在包目录中，所有以_test.go为后缀的源码文件都会被go test运行到

go build命令不会将这些测试文件打包到最后的可执行文件中，**测试用例与被测试的文件必须在同目录下**

```go
// /ch07/coding/add_test.go
func TestAdd(t *testing.T) {
	res := add(1, 2)

	if res != 4 {
		t.Errorf("错误")
	}
}

// /ch07/coding/add.go
func add(a, b int) int {
	return a + b
}

```

## 2024/5/03 Day7

leetcode：

[有效的数独](https://leetcode.cn/problems/valid-sudoku/)
