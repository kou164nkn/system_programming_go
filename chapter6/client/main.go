package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil

	// enclose in loop for retry
	for {
		var err error

		// no connection yet / retry on error
		if conn == nil {
			// initialize conn from Dial
			conn, err = net.Dial("tcp", "localhost:8080")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Accept %d\n", current)
		}

		// create POST request
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8080",
			strings.NewReader(sendMessages[current]),
		)
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		// read from server
		// Timeout cause an error here, so retry
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		// view result
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		// if all transmission is completed, it ends
		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
