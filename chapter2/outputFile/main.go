package main

import (
	"fmt"
	"os"
)

func createFileAndContent(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	message := "Hello World!"
	intNum := 5
	floatNum := 3.14

	fmt.Fprintf(file, "print message: %v\nprint decimal number: %d\nprint float number:%v\n", message, intNum, floatNum)
	return nil
}

func main() {
	fileName := "output.txt"

	if err := createFileAndContent(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
