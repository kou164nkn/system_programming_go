package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/shirou/gopsutil/process"
)

func main() {
	path, _ := os.Executable()
	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("実行ファイルパス: %s\n", path)

	fmt.Printf("プロセスID: %d\n", os.Getegid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())

	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Printf("グループID: %d\n", syscall.Getpgrp())
	fmt.Printf("セッションID: %d\n", sid)

	fmt.Printf("ユーザーID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getuid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %v\n", groups)

	fmt.Printf("実効ユーザーID: %d\n", os.Geteuid())
	fmt.Printf("実効グループID: %d\n", os.Getegid())

	wd, _ := os.Getwd()
	fmt.Printf("作業ディレクトリ: %s\n", wd)

	fmt.Println(os.ExpandEnv("${HOME}/gobin"))

	p, _ := process.NewProcess(int32(os.Getppid()))
	name, _ := p.Name()
	cmd, _ := p.Cmdline()
	fmt.Printf("parent pid: %d\n", p.Pid)
	fmt.Printf("parent name: %s\n", name)
	fmt.Printf("parent cmd: %s\n", cmd)
}
