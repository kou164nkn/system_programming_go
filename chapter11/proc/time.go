package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	state := cmd.ProcessState
	fmt.Printf("%s\n", state.String())
	fmt.Printf("  Pid: %d\n", state.Pid())
	fmt.Printf("  Exited: %v\n", state.Exited())
	fmt.Printf("  ExitCode: %v\n", state.ExitCode())
	fmt.Printf("  Success: %v\n", state.Success())
	// システム時間(カーネル内で行われた処理の時間)
	fmt.Printf("  System: %v\n", state.SystemTime())
	// ユーザー時間(プロセス内で消費された時間)
	fmt.Printf("  User: %v\n", state.UserTime())
}
