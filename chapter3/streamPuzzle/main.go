package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	getA := io.NewSectionReader(programming, 5, 1)
	getS := io.LimitReader(system, int64(1))
	getC := io.LimitReader(computer, int64(1))

	sr := io.NewSectionReader(programming, 8, 1)

	pr, pw := io.Pipe()
  mw := io.MultiWriter(pw, pw)

	go io.CopyN(mw, sr, 2)
  defer pw.Close()

  // io.LimitReader(pr, 2)とすることでprは2度入力待ち状態になる
	stream = io.MultiReader(getA, getS, getC, io.LimitReader(pr, 2))
	io.Copy(os.Stdout, stream)
}
