package main

import (
	"net/rpc"
	"jsj.golangtc/study01"
	"net"
	"net/rpc/jsonrpc"
	"log"
)

func main() {

	rpc.Register(study01.DemoService{})

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error : %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

}
