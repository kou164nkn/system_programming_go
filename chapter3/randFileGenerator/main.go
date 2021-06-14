package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("rand.bin")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
	}
	defer file.Close()

  var binSize int64 = 1024
	lr := io.LimitReader(rand.Reader, binSize)
	if _, err := io.CopyN(file, lr, binSize); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

	// DEBUG: check file size
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	fmt.Println(size)
}
