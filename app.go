package main

import (
	"flag"
	"./node"
)

const (
	defaultHost   = "127.0.0.1"
	defaultPort   = 8080
)


func main() {
	flag.String("host", defaultHost, "default host")
	flag.Int("port", defaultPort, "default port")

	master := node.GetMaster()
	master.Serve()
}