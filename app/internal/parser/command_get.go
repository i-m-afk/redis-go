package parser

import (
	"fmt"
	"log"

	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
)

func commandGet(conf *configuration.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing key in args")
	}

	key := args[0]

	conf.Mux.Lock()
	defer conf.Mux.Unlock()

	value, ok := conf.Store[key]
	if !ok {
		conf.Conn.Write([]byte("$-1\r\n"))
		return fmt.Errorf("key %s not found", key)
	}
	// TODO:
	// format (hack)
	respValue := fmt.Sprintf("$%d\r\n%s\r\n", len(value.Value), value.Value)

	_, err := conf.Conn.Write([]byte(respValue))
	if err != nil {
		log.Println("Couldn't write here")
		return err
	}

	return nil
}
