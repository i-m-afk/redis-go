package parser

import (
	"log"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
)

func commandSet(conf *configuration.Config, args ...string) error {
	key := args[0]
	value := args[1]

	store := configuration.RedisStore{}
	if len(args) > 2 {
		ttl, err := strconv.ParseInt(args[2], 10, 32)
		if err != nil {
			log.Println("Failed to parse int")
		}
		store = configuration.RedisStore{Value: value, Type: configuration.BulkString, Ttl: int(ttl)}
	} else {
		store = configuration.RedisStore{Value: value, Type: configuration.BulkString, Ttl: 10000000}
	}

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
