package main

import (
	"flag"
	"fmt"
	"github.com/quickfixgo/quickfix"
	"os"
	"os/signal"
	"path"
)

var SEP = "#"
var version = "1.0.6 - Generate-Fix"

func main() {
	flag.Parse()

	/**
	Creating input and output channels for goroutines communication
	*/
	proxyChan := make(chan quickfix.Messagable)
	serverChan := make(chan quickfix.Messagable)

	/**
	Creating FIX4.4 Server that will receive connections
	*/
	srvCfgFileName := path.Join("config", "cache_server.cfg")
	if flag.NArg() > 0 {
		srvCfgFileName = flag.Arg(0)
	}
	cacheServer := newCacheServer(serverChan, proxyChan)
	go cacheServer.Start(srvCfgFileName)

	/**
	Creating FIX4.4 Initiator that will connect to B3
	*/
	clientCfgFileName := path.Join("config", "cache_proxy.cfg")
	if flag.NArg() > 1 {
		clientCfgFileName = flag.Arg(1)
	}
	cacheProxy := newCacheProxy(proxyChan, serverChan)
	go cacheProxy.Start(clientCfgFileName)

	fmt.Println("CACHE VERSION: " + version)
	// Wait for interrupt
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt

	fmt.Println("Gonna stop server")
	cacheServer.Stop()
	//cacheProxy.Stop()

	fmt.Println("Gonna close channels")
	close(proxyChan)
	close(serverChan)

	fmt.Println("DONE.")
}
