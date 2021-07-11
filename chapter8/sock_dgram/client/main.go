package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	os.Remove(clientPath)

	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	var serverAddr net.Addr = unixServerAddr
	if err != nil {
		panic(err)
	}

	log.Println("Sending to server")
	_, err = conn.WriteTo([]byte("Hello from client"), serverAddr)
	if err != nil {
		panic(err)
	}

	log.Println("Receive from server")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received %v: %v\n", serverAddr, string(buffer[:length]))
}
