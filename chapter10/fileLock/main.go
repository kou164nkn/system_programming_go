// +build darwin dragonfly freebsd linux netbsd openbsd

package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

type FileLock struct {
	l  sync.Mutex
	fd int
}

func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic("filename needed")
	}
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, 0750)
	if err != nil {
		panic(err)
	}

	return &FileLock{fd: fd}
}

func (fl *FileLock) Lock() {
	fl.l.Lock()
	if err := syscall.Flock(fl.fd, syscall.LOCK_EX); err != nil {
		panic(err)
	}
}

func (fl *FileLock) Unlock() {
	if err := syscall.Flock(fl.fd, syscall.LOCK_UN); err != nil {
		panic(err)
	}
	fl.l.Unlock()
}

func main() {
	fl := NewFileLock("main.go")

	fmt.Println("try locking...")
	fl.Lock()
	fmt.Println("main.go is locked")

	time.Sleep(10 * time.Second)

	fl.Unlock()
	fmt.Println("main.go is unlocked")
}
