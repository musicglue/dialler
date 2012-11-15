package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var addr = flag.String("address", "127.0.0.1:80", "A hostname or IP with an optional port")
var timeout = flag.Int("timeout", 1, "Timeout in millieconds")

func main() {
	flag.Parse()
	_, err := net.DialTimeout("tcp", *addr, time.Duration(*timeout)*time.Millisecond)
	if err != nil {
		log.Fatal("Not Connected")
	}
	fmt.Println("Connected to Server")
}
