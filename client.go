package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}

	if resp != "OK\n" {
		conn.Close()
		log.Fatal("invalid response")
	}
}
