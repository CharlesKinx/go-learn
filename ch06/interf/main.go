package main

import "fmt"

/*
*
定义接口, 接口只有定义，必须使用结构体将所有的方法实现
*/

type Duck interface {
	// 接口里面主要是方法

	Gaga()
	Walk()
	Swimming()
}

/*
*
结构体
*/
type pskDuck struct {
	name string
	age  int
}

/*
*
结构体实现接口的方法
*/
func (pd *pskDuck) Gaga() {
	fmt.Println("gagaga...")
}

func (pd *pskDuck) Walk() {
	fmt.Println("walk...")
}

func (pd *pskDuck) Swimming() {
	fmt.Println("Swimming...")
}

func main() {

	var duck Duck = &pskDuck{
		name: "pskDuck",
		age:  12,
	}

	duck.Walk()

}
