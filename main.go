package main

import (
	"flag"
	"log"

	"github.com/chlins/boring/client"
	"github.com/chlins/boring/server"
)

var (
	s = flag.Bool("s", false, "server mode")
	c = flag.Bool("c", false, "client mode")
	a = flag.String("a", "", "server address")
)

func main() {
	flag.Parse()

	if *s && *c {
		flag.Usage()
		return
	}

	if *s {
		sver := server.New()
		log.Printf("Listen address: %s\n", sver.Addr())
		sver.Accept()
	}

	if *c {
		client.New(*a).RWLoop()
	}
}
