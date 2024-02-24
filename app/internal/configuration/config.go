package configuration

import (
	"net"
	"sync"
)

type Config struct {
	Conn  net.Conn
	Store map[string]RedisStore
	Mux   sync.Mutex
}

type RedisStore struct {
	Value string
	Type  RESPType
}

type RESPType byte

const (
	SimpleString RESPType = iota
	Error                 = '-'
	Integer               = ':'
	BulkString            = '$'
	String                = '+'
	Array                 = '*'
)

func (conf *Config) InitConfig(conn net.Conn) *Config {
	conf.Conn = conn
	conf.Store = make(map[string]RedisStore)
	conf.Mux = sync.Mutex{}
	return conf
}
