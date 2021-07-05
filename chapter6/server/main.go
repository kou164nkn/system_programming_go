package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// loop to respond multiple times on the socket after accept
			for {
				// setting up timeout
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				// reading http request
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// exit when timeout or socket closed
					// anything else is error
					neterr, ok := err.(net.Error) // downcast
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				// view request
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))

				content := "Hello World\n"

				// writing http request
				// HTTP/1.1 and contentlength must be set
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}
