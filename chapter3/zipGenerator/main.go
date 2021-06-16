package main

import (
	"archive/zip"
	"log"
	"os"
)

// 問題文の意図が掴めなかったので後回し
func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
}
