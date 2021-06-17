package main

import (
	"archive/zip"
	"log"
	"net/http"
)

func ZipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	zw := zip.NewWriter(w)
	defer zw.Close()

	zipContents := [...]struct {
		name, body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, content := range zipContents {
		f, err := zw.Create(content.name)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = f.Write([]byte(content.body))
		if err != nil {
			log.Print(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", ZipHandler)
	http.ListenAndServe(":8080", nil)
}
