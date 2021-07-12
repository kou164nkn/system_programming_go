package main

import (
	"fmt"
	"os"
	"syscall"
)

func printFileInfo(fileName string) {
	info, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("FileInfo")
	fmt.Printf("  ファイル名: %v\n", info.Name())
	fmt.Printf("  サイズ: %v\n", info.Size())
	fmt.Printf("  変更日時: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("  ディレクトリ?  %v\n", info.Mode().IsDir())
	fmt.Printf("  読み書き可能な通常ファイル?  %v\n", info.Mode().IsRegular())
	fmt.Printf("  Unixのファイルアクセス権限ビット  %o\n", info.Mode().Perm())
	fmt.Printf("  モードのテキスト表現  %v\n", info.Mode().String())

	internalStat := info.Sys().(*syscall.Stat_t)
	fmt.Println("Sys()")
	fmt.Printf("  デバイス番号: %v\n", internalStat.Dev)
	fmt.Printf("  inode番号: %v\n", internalStat.Ino)
	fmt.Printf("  ブロックサイズ: %v\n", internalStat.Blksize)
	fmt.Printf("  ブロック数: %v\n", internalStat.Blocks)
	// error occured: internalStat.NLink undefined (type *syscall.Stat_t has no field or method NLink, but does have Nlink)
	// fmt.Printf("  リンクされている数: %v\n", internalStat.NLink)
	fmt.Printf("  ファイル作成日時: %v\n", internalStat.Birthtimespec)
	fmt.Printf("  最終アクセス日時: %v\n", internalStat.Atimespec)
	fmt.Printf("  属性変更日時: %v\n", internalStat.Ctimespec)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}

	printFileInfo(os.Args[0])
}
