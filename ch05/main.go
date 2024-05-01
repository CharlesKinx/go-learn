package main

import "fmt"

type Person struct {
	name string
}

/*
*
普通传递
*/
func changename(p Person) {
	p.name = "hello"
}

/*
*
指针传递
*/
func changename1(p *Person) {
	p.name = "point"
}

func (p *Person) sayHello() {
	p.name = "ppp"
}

func main() {

	p := Person{
		name: "test",
	}

	fmt.Println(p)
	changename(p)
	fmt.Println(p)
	changename1(&p)
	fmt.Println(p)

	var a int = 10
	b := &a
	fmt.Println(*b)

	// 指针一定要初始化,第一种初始化方法
	ps := &Person{}
	ps.name = "ps"
	// 第二种初始化方法
	var ps1 Person
	ps2 := ps1
	ps2.name = "ps2"
	// 第三种初始化方法
	var ps3 *Person = new(Person)
	ps3.name = "ps3"
	fmt.Println(ps3)
	ps3.sayHello()
	fmt.Println(ps3)
	// 指针一定要初始化，负责会出现nil pointer
	// map也一定要初始化
	//hash := make(map[string]string)
	//hash["haha"] = "value"

}
