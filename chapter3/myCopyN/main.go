package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// io.Copy is used, so the buffer allocated is 32KB.
func myCopyN(dest io.Writer, src io.Reader, length int) (int64, error) {
	ltdReader := io.LimitReader(src, int64(length))
	return io.Copy(dest, ltdReader)
}

func main() {
	var buf bytes.Buffer

	str := "Hello, World!"
	br := bytes.NewBuffer([]byte(str))

	n, err := myCopyN(&buf, br, 10)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Printf("read size: %d\n", n)
	fmt.Println(buf.String())
}
