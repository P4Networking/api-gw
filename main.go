package main

import (
	"github.com/P4Networking/api-gw/server"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go server.StartGRPCServer(&wg)
	go server.StartHttpReverseProxyServer(&wg)

	wg.Wait()
}
