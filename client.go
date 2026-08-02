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
	defer conn.Close()

	resp, _ := bufio.NewReader(conn).ReadString('\n')
	if resp != "OK\n" {
		log.Fatal("invalid response")
	}
}
