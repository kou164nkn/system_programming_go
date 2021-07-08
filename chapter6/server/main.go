package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// export to conn in order (execute in goroutine)
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()

	for sessionResponse := range sessionResponses {
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

// handle the session in the request
func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	content := "Hello, World\n"
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          io.NopCloser(strings.NewReader(content)),
	}
	resultReceiver <- response
}

// procssing 1 session
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	// channel for processing requests in a session in order
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)

	// goroutine that serializes the response and writes it to the socket
	go writeToConn(sessionResponses, conn)

	reader := bufio.NewReader(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// reading http request
		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			}
			if err == io.EOF {
				break
			}
			panic(err)
		}

		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse

		// return response asynchronously
		go handleRequest(request, sessionResponse)
	}
}

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

		go processSession(conn)
	}
}
