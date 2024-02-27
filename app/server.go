package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/internal/configuration"
	"github.com/codecrafters-io/redis-starter-go/app/internal/parser"
)

func main() {
	// configs
	dir := flag.String("dir", "", "directory where RDB files are stored")
	dbfilename := flag.String("dbfilename", "", "Name of the RDB file")
	flag.Parse()
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("SERVER LISTENING ON 0.0.0.0:6379")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		conf := configuration.Config{}
		conf.InitConfig(conn, dir, dbfilename)
		go handleConnection(&conf)
	}
}

func handleConnection(conf *configuration.Config) (err error) {
	conn := conf.Conn
	defer conn.Close()
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("Error reading data from connection: ", err.Error())
			return err
		}
		parser.Parse(conf, string(buffer[:n]))
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
