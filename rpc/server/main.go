package main

import (
	"net/rpc"
	"jsj.golangtc/rpc"
	"net"
	"net/rpc/jsonrpc"
	"log"
)

func main() {

	rpc.Register(rpcdemo.DemoService{})
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
	//可用telnet后,输入
	//{"method":"DemoService.Div", "params":[{"A":3, "B":2}], "id":123}


}
