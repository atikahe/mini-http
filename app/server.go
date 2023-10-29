package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening", l.Addr())

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Request accepted")

	buffer := make([]byte, 4096)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading request content", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	if err != nil {
		fmt.Println("Error sending response", err.Error())
		os.Exit(1)
	}
}
