package parser

import (
	"fmt"
	"net"
)

func ParseCommand(conn net.Conn, commandName string, args []string) {
	command, exits := Commands()[commandName]
	if exits {
		err := command.callback(conn, args...)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Unknown command %v\n", commandName)
	}
}
