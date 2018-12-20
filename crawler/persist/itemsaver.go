package persist

import (
	"fmt"
	"jsj.golangtc/crawler/zhenai/engine"
)

func ItemSaver() chan engine.Item {

	out := make(chan engine.Item)

	itemCount := 0
	go func() {
		for {
			item := <-out
			fmt.Printf("%s: %v\n", item.Url, item.Payload)
			itemCount ++
		}
	}()

	return out
}
