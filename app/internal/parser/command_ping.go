package parser

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
)

func commandPing(conf *configuration.Config, args ...string) error {
	conn := conf.Conn
	_, err := conn.Write([]byte(fmt.Sprintf("+PONG\r\n")))
	if err != nil {
		fmt.Println("Error occured in writing ", err)
		return err
	}
	return nil
}
