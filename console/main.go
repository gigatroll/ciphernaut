package main

import (
	".."
	"fmt"
	"strconv"
)

var (
	terminateProgram chan bool
)

func main() {
	addr := "127.0.0.1"
	port := 9999

	fmt.Println(addr + ":" + strconv.Itoa(port))
	ciphernaut.StartTCPProxy(addr, port)

	for {
		select {
		case <-terminateProgram:
			return
		}
	}
}
