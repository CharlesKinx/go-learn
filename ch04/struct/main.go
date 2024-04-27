package main

import "fmt"

// 结构体嵌套，类似于子类复用父类的属性 子结构体
type Student struct {
	// 第一种嵌套方式
	p Person

	// 第二种嵌套方式，匿名嵌套
	// Person
	score float32
}

// 结构体
type Person struct {
	name    string
	age     int
	address string
	height  float32
}

// 结构体绑定方法,接收器有两种形态 该函数为值传递，传递的是Person的拷贝
// 另外一种为指针传递， p *Person
func (p Person) fun_person() {
	fmt.Println(p)
}

func main() {

	p1 := Person{
		name:    "Whd",
		age:     24,
		address: "hangdian",
		height:  1.80,
	}

	fmt.Println(p1)

	var persons []Person

	persons = append(persons, p1)
	persons = append(persons, Person{
		name:    "Whd1a",
		age:     24,
		address: "hangdian1a",
		height:  1.80,
	})

	fmt.Println(persons)

	// 匿名结构体

	address := struct {
		province string
		city     string
		address  string
	}{
		province: "北京",
		city:     "beijing",
		address:  "xxx",
	}

	fmt.Println(address)

	// 结构体嵌套
	student := Student{
		p:     p1,
		score: 98.5,
	}

	student2 := Student{
		p: Person{
			name:    "吴海董",
			age:     24,
			address: "杭州",
			height:  1.80,
		},
		score: 97,
	}

	fmt.Println(student)
	fmt.Println(student2)
	// 调用结构体方法
	p1.fun_person()
	student.p.fun_person()
}
