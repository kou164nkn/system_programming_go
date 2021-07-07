package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create POST request
	request, err := http.NewRequest(
		"POST",
		"http://localhost:8080",
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	// read from server
	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}

	// view result
	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		// obtain size
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		// parse hexadecimal format size
		// close if size is zero
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 24)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		// allocate a buffer for the number of sizes and read
		line := make([]byte, int(size))
		io.ReadFull(reader, line)
		reader.Discard(2)
		fmt.Printf("  %d bytes: %s\n", size, string(line))
	}
}
