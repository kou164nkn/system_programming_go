package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

func main() {
	// reading testdata
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := os.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}

	// mapping to memory
	// m is alias of []byte, so m is able to access with index
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer m.Unmap()

	// fix data on the memory
	m[9] = 'X'
	m.Flush()

	// try to read
	fileData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap:     %s\n", m)
	fmt.Printf("file:     %s\n", fileData)
}
