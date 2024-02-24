package parser

import (
	"fmt"
	"net"
)

func commandPong(conn net.Conn, args ...string) error {
	_, err := conn.Write([]byte(fmt.Sprintf("+PONG\r\n")))
	if err != nil {
		fmt.Println("Error occured in writing ", err)
		return err
	}
	return nil
}
