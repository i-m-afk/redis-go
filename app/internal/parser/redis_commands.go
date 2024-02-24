package parser

import "net"

type RedisCommand struct {
	name        string
	description string
	callback    func(net.Conn, ...string) error
}

func Commands() map[string]RedisCommand {
	return map[string]RedisCommand{
		"ping": {
			name:        "ping",
			description: "Returns PONG if no argument is provided, otherwise return a copy of the argument as a bulk.",
			callback:    commandPong,
		},
		"echo": {
			name:        "echo",
			description: "Returns message.",
			callback:    commandEcho,
		},
	}
}
