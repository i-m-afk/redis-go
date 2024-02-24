package parser

import (
	"log"

	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
)

func commandSet(conf *configuration.Config, args ...string) error {
	key := args[0]
	value := args[1]
	store := configuration.RedisStore{Value: value, Type: configuration.BulkString}

	conf.Mux.Lock()
	defer conf.Mux.Unlock()
	conf.Store[key] = store
	_, err := conf.Conn.Write([]byte("+OK\r\n"))
	if err != nil {
		log.Println("failed to write")
		return err
	}
	return nil
}
