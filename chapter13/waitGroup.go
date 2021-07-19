package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// ジョブ数をあらかじめ登録
	wg.Add(2)

	go func() {
		fmt.Println("Work1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Work2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finished")
}
