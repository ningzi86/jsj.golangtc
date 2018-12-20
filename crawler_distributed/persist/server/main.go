package main

import (
	"jsj.golangtc/crawler_distributed/persist"
	"jsj.golangtc/crawler_distributed/rpcsupport"
)

func main() {
	ServerRpc(":1234")
}

func ServerRpc(host string) {
	rpcsupport.ServerRPC(host, persist.ItemSaverService{})
}
