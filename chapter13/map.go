package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := sync.Map{}

	// 任意の型の値を格納できる
	smap.Store("hello", "world")
	smap.Store(1, 2)

	// キーを指定して削除
	smap.Store("test", "yeah")
	smap.Delete("test")

	fmt.Println(smap)

	// 取り出し方法
	value, ok := smap.Load("hello")
	fmt.Printf("key: %v, value: %v, exists?: %v\n", "hello", value, ok)

	// キーの有無で、登録の可否を判断
	// 前者はすでにキーが存在するので無視
	smap.LoadOrStore(1, 3)
	smap.LoadOrStore(2, 4)

	// rangeに対応していない代わりに、ループ用のメソッドを持つ
	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
}
