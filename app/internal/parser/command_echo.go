package parser

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
)

func commandEcho(conf *configuration.Config, args ...string) error {
	conn := conf.Conn
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
