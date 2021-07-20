package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

func main() {
	cmd := exec.Command("./check")
	stdpty, stdtty, _ := pty.Open()
	defer stdtty.Close()

	cmd.Stdin = stdpty
	cmd.Stdout = stdpty

	errpty, errtty, _ := pty.Open()
	defer errtty.Close()

	cmd.Stderr = errpty

	go func() {
		io.Copy(os.Stdout, stdpty)
	}()
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}