package client

import (
	"jsj.golangtc/crawler/zhenai/engine"
	"jsj.golangtc/crawler_distributed/rpcsupport"
	"fmt"
)

func ItemSaver() chan engine.Item {

	out := make(chan engine.Item)
	client, _ := rpcsupport.NewClinet(":1234")

	itemCount := 0
	go func() {
		for {
			item := <-out
			var result string
			client.Call("ItemSaverService.Save", item, &result)

			fmt.Println(result)
			//fmt.Printf("%s: %v\n", item.Url, item.Payload)
			itemCount ++
		}
	}()

	return out
}
