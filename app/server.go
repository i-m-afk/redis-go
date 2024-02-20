package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("SERVER LISTENING ON 0.0.0.0:6379")

	for {

		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// TODO: Read data (write a command parser)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data from connection: ", err.Error())
		return
	}

	fmt.Println("Recieved : ", string(buffer[:n]))

	conn.Write([]byte(fmt.Sprintf("+PONG\r\n")))
	conn.Write([]byte(fmt.Sprintf("+PONG\r\n")))
}
