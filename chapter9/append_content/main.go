package main

import (
	"io"
	"os"
)

func append() {
	// OpenFile() returned error when textfile.txt doesn't exist.
	//
	// os.Create() called OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	// O_TRUNC indicate truncate regular writable file when opened.
	//
	// O_APPEND indicate append data to the file when writing.
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.WriteString(file, "Append content\n")
}

func main() {
	append()
}
