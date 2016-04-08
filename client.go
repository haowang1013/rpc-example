package main

import (
	"github.com/haowang1013/rpc-example/shared"
	"github.com/op/go-logging"
	"net"
	"net/rpc"
	"time"
)

var log = logging.MustGetLogger("")

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:10000", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	var ret shared.CacheItem
	err = client.Call("Server.Echo", shared.CacheItem{"Test", "Test"}, &ret)
	if err != nil {
		log.Error(err)
	} else {
		log.Debugf("Got response: %v", ret)
	}
}
