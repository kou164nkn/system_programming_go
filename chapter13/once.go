package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("Initializing...")
}

var once sync.Once

func main() {
	// 3回呼び出しても一度しか実行されない
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
