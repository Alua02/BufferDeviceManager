package main

import (
	"fmt"
	"net"
)

func TcpAndroid() {
	listener, err := net.Listen("tcp", "7777")
	if err != nil {
		fmt.Println("error listening on port 7777")
	}
	
}
