package main

import (
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:7")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 100)

	for {
		sizeRead, err := conn.Read(buffer)
		if err != nil {
			return
		}
		sizeWritten, err := conn.Write(buffer[:sizeRead])
		if err != nil {
			return
		}
		for sizeWritten != sizeRead {
			moreWritten, err := conn.Write(buffer[sizeWritten:sizeRead])
			if err != nil {
				return
			}
			sizeWritten += moreWritten
		}
	}
}
