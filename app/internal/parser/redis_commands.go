package parser

import (
	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
)

type RedisCommand struct {
	name        string
	description string
	callback    func(*configuration.Config, ...string) error
}

func Commands() map[string]RedisCommand {
	return map[string]RedisCommand{
		"ping": {
			name:        "ping",
			description: "Returns PONG if no argument is provided, otherwise return a copy of the argument as a bulk.",
			callback:    commandPing,
		},
		"echo": {
			name:        "echo",
			description: "Returns message.",
			callback:    commandEcho,
		},
		"set": {
			name:        "set",
			description: "Set key to hold the string value. If key already holds a value, it is overwritten, regardless of its type. Any previous time to live associated with the key is discarded on successful SET operation.",
			callback:    commandSet,
		},
		"get": {
			name:        "get",
			description: "Get the value of key. If the key does not exist the special value nil is returned. An error is returned if the value stored at key is not a string, because GET only handles string values.",
			callback:    commandGet,
		},
	}
}
