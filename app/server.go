package main

import (
	"errors"
	"fmt"
	"io"
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) (err error) {
	for {
		// TODO: Read data (write a command parser)
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("Error reading data from connection: ", err.Error())
			return err
		}

		fmt.Println("Recieved : ", string(buffer[:n]))

		conn.Write([]byte(fmt.Sprintf("+PONG\r\n")))
	}

	return nil
}
