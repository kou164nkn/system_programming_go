package main

import (
	"flag"
	"fmt"
  "os"

	"github.com/kou164nkn/system_programming_go/chapter3/fileCopy"
)

var fileName, newFileName string

func init() {
	flag.StringVar(&fileName, "src", "oldFile.txt", "source file name")
	flag.StringVar(&newFileName, "dst", "newFile.txt", "new file name")
}

func main() {
	flag.Parse()

	if err := fileCopy.Copy(fileName, newFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
