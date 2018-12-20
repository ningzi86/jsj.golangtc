package main

import (
	"testing"
	"jsj.golangtc/crawler/zhenai/engine"
	"jsj.golangtc/crawler_distributed/rpcsupport"
	"fmt"
	"time"
)

func TestItemServer(t *testing.T) {

	host := ":1234"
	go ServerRpc(host)

	time.Sleep(time.Second)

	item := &engine.Item{
		Url:  "http://www.baidu.com",
		Type: "type",
		Id:   "Id",
	}

	client, err := rpcsupport.NewClinet(":1234")

	if err != nil {
		panic(err)
	}

	var result string
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil {
		panic(err)
	}

	fmt.Println("result", result)

}
