package main

import (
	"fmt"
	"sync"
)

// 锁的本质是将并行的代码串行化了，使用lock肯定会影响性能

// 即使是设计锁，那么也应该尽量保证并行

// 现在有两组协程，一组负责写数据，一组负责度数据

func main() {

	var rwlock sync.RWMutex
	var total int
	var wg sync.WaitGroup

	wg.Add(2)
	// 写goroutine
	go func() {
		defer wg.Done()
		// 加写锁，写锁会防止别的写锁以及读锁
		rwlock.Lock()
		defer rwlock.Unlock()

		total = 12
	}()

	go func() {
		defer wg.Done()
		// 加读锁，读锁不会阻止别的读锁
		rwlock.RLock()
		defer rwlock.RUnlock()
		fmt.Println(total)
	}()

	wg.Wait()
}
