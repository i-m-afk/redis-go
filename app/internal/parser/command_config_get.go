package parser

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
)

func commandConfigGet(conf *configuration.Config, args ...string) error {
	if args[0] != "get" {
		return fmt.Errorf("Unknown command config %s", args[0])
	}

	var respString string

	if args[1] == "dir" {
		respString = fmt.Sprintf("*2\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(args[1]), args[1], len(conf.Dir), conf.Dir)
	} else if args[1] == "dbfilename" {
		respString = fmt.Sprintf("*2\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(args[1]), args[1], len(conf.DbName), conf.DbName)
	} else {
		return fmt.Errorf("Unknown config option %s", args[1])
	}

	_, err := conf.Conn.Write([]byte(respString))
	if err != nil {
		return err
	}

	return nil
}
