package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine 如何使用

func go_routine() {
	for i := 0; i < 100; i++ {

		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(10 * time.Second)
}

// 子goroutine如何通知主goroutine自己执行结束
/**
WaitGroup 主要用于goroutine的执行等待，Add方法要和Done结合
*/
func go_routine_wait() {
	var ws sync.WaitGroup
	// 需要监控多少个goroutine结束
	ws.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {

			// defer ws.Done() 也可以使用defer来
			fmt.Println(i)
			// 结束
			ws.Done()
		}(i)
	}

	ws.Wait()
}
func main() {

	go_routine_wait()
}
