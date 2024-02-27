package configuration

import (
	"net"
	"sync"
	"time"
)

type Config struct {
	Conn   net.Conn
	Store  map[string]RedisStore
	Mux    sync.Mutex
	Dir    string
	DbName string
}

type RedisStore struct {
	Value string
	Type  RESPType
	Ttl   int
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

func (conf *Config) InitConfig(conn net.Conn, dir *string, dbfilename *string) *Config {
	conf.Conn = conn
	conf.Dir = *dir
	conf.DbName = *dbfilename
	conf.Store = make(map[string]RedisStore)
	go conf.reapLoop(1 * time.Millisecond)
	return conf
}

func (conf *Config) reapLoop(dur time.Duration) {
	tick := time.NewTicker(dur)
	defer tick.Stop()
	for range tick.C {
		conf.reap(dur)
	}
}

func (conf *Config) reap(dur time.Duration) {
	conf.Mux.Lock()
	defer conf.Mux.Unlock()
	currentTimestamp := time.Now().Unix()
	for key, val := range conf.Store {
		if int64(val.Ttl) < currentTimestamp {
			delete(conf.Store, key)
		}
	}
}
