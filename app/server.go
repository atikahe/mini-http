package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	request, err := DecodeRequest(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statusCode := StatusNotFound
	if request.Path == "/" {
		statusCode = StatusOK
	}

	response := fmt.Sprintf("HTTP/1.1 %d %s OK\r\n\r\n", statusCode, statusCode)

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response", err.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening", l.Addr())

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Client connected")

		go handleConnection(conn)
	}
}
