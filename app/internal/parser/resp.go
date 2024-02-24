package parser

import (
	"net"
	"strings"
)

type RESPType int

const (
	SimpleString RESPType = iota
	Error                 = '-'
	Integer               = ':'
	BulkString            = '$'
	String                = '+'
	Array                 = '*'
)

func compare(ch byte) bool {
	return ch == Error || ch == Integer || ch == BulkString || ch == String || ch == Array
}
func Parse(conn net.Conn, input string) {
	parts := strings.Split(strings.ToLower(input), "\r\n")
	// hack
	something := make([]string, 0)
	for i, part := range parts {
		if len(part) == 0 || part == "\n" {
			continue
		}
		if !compare(part[0]) {
			something = append(something, parts[i])
		}
	}
	command := something[0]
	ParseCommand(conn, command, something[1:])
}
