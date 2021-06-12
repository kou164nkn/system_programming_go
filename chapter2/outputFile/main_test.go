package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestCreateFileAndContent(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		fileName      string
		expectContent string
	}{
		"success": {fileName: "test.txt", expectContent: "print message: Hello World!\nprint decimal number: 5\nprint float number:3.14\n"},
	}

	for name, tt := range cases {
		tt := tt

		t.Run(name, func(t *testing.T) {
			createFileAndContent(tt.fileName)

			file, err := os.Open(tt.fileName)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			fileInfo, _ := file.Stat()
			var size int64 = fileInfo.Size()
			actualContent := make([]byte, size)

			buffer := bufio.NewReader(file)
			_, err = buffer.Read(actualContent)

			if !reflect.DeepEqual(string(actualContent), tt.expectContent) {
				t.Errorf("createFileAndContent want %v but %v", tt.expectContent, actualContent)
			}
		})
		flushTestResult(t, tt.fileName)
	}
}

func flushTestResult(t *testing.T, fileName string) {
	t.Helper()

	if err := os.Remove(fileName); err != nil {
		t.Fatal(err)
	}
}
