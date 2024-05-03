package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
go 语言互斥锁的用法

用来资源竞争
*/

var total int32
var wg sync.WaitGroup

// 锁不能被复制
var lc sync.Mutex

func add() {

	atomic.AddInt32(&total, 1)
	lc.Lock()
	for i := 0; i < 100000; i++ {
		total += 1
	}
	lc.Unlock()
	wg.Done()
}

func sub() {
	atomic.AddInt32(&total, -1)

	//lc.Lock()
	//for i := 0; i < 100000; i++ {
	//	total -= 1
	//}
	//lc.Unlock()
	wg.Done()
}

func main() {

	wg.Add(2)
	go add()
	go sub()
	wg.Wait()

	fmt.Println(total)
}
