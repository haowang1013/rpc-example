package main

import (
	"github.com/haowang1013/rpc-example/shared"
	"github.com/op/go-logging"
	"net"
	"net/rpc"
)

var log = logging.MustGetLogger("")

type Server struct{}

func (s *Server) Echo(arg shared.CacheItem, ret *shared.CacheItem) error {
	log.Debugf("Server.Echo: %v", arg)
	*ret = arg
	return nil
}

func main() {
	server := &Server{}
	rpc.Register(server)

	l, e := net.Listen("tcp", ":10000")
	if e != nil {
		log.Fatal(e)
	}
	defer l.Close()

	rpc.Accept(l)
}
