package main

import (
	"net"
	"net/rpc/jsonrpc"
	"jsj.golangtc/rpc"
	"fmt"
)

func main() {

	conn, err := net.Dial("tcp", ":1234")
	if (err != nil) {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{A: 100, B: 22}, &result)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
