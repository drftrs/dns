package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
)

var directives = []string{
	"drftrs",
	"cache",
	"log",
	"forward",
	"startup",
	"shutdown",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	go coremain.Run()
	select {}
}
