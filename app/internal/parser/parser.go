package parser

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
)

func ParseCommand(conf *configuration.Config, commandName string, args []string) {
	command, exits := Commands()[commandName]
	if exits {
		err := command.callback(conf, args...)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Unknown command %v\n", commandName)
	}
}
