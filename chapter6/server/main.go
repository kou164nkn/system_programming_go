package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// 青空文庫: ごんぎつねより
// http://www.aozora.gr.jp/cards/000121/card628.html
var contents = []string{
	"これは、私わたしが小さいときに、村の茂平もへいというおじさんからきいたお話です。",
	"むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって",
	"中山さまというおとのさまが、おられたそうです。",
	"その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。",
	"ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

// procssing 1 session
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		// reading http request
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
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

		// writing http response
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
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
