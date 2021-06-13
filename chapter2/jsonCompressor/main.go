package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// write the test later
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
    "これは": "JSONをZIP化するサーバ",
	}

	gw := gzip.NewWriter(w)
	defer gw.Close()
	mw := io.MultiWriter(gw, os.Stdout)

	en := json.NewEncoder(mw)
	err := en.Encode(source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Encode error: ", err)
	}

	if err = gw.Flush(); err != nil {
		fmt.Fprintln(os.Stderr, "Flush error: ", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
