package main

import (
	"net"
	"net/rpc/jsonrpc"
	"fmt"
	"jsj.golangtc/study01"
)

func main() {

	conn, err := net.Dial("tcp", ":1234")
	if (err != nil) {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float32
	err = client.Call("DemoService.Div",study01.Args{Number1: 100, Number2: 22}, &result)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
