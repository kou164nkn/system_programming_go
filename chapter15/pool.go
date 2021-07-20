package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int

	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	pool.Put("manualy added: 1")
	pool.Put("manualy added: 2")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get()) // 新規作成

	pool.Put("manualy added: 3")
	runtime.GC()
	fmt.Println(pool.Get()) // GCを呼ぶと追加した要素が消える
}
