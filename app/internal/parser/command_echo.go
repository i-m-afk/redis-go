package parser

import (
	"fmt"
	"net"
	"strings"
)

func commandEcho(conn net.Conn, args ...string) error {
	var str strings.Builder
	for _, arg := range args {
		l := len(arg)
		str.WriteString(fmt.Sprintf("$%d\r\n", l))
		str.WriteString(fmt.Sprintf("%s\r\n", arg))
	}
	_, err := conn.Write([]byte(str.String()))
	if err != nil {
		return err
	}
	return nil
}
