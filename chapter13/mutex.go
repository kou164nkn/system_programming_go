package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	// Lock()/Unlock()のペアを呼び出してロックする
	mutex.Lock()
	defer mutex.Unlock()
	id++
	result := id
	return result
}

func main() {
	maxCount := 20

	// sync.Mutex構造体の変数宣言
	// 次の宣言をしてもポインタ型になるだけで正常に動作する
	// mutex := sync.Mutex
	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(maxCount)

	for i := 0; i < maxCount; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
			(&wg).Done()
		}()
	}

	wg.Wait()
}
