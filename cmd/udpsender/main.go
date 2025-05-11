package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	address = ":42069"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalf("unable to resolve UDP address: %v", err)
	}
	con, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("unable to dial UDP: %v", err)
	}
	defer con.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("readstring err: %v", err)
		}
		if _, err := con.Write([]byte(line)); err != nil {
			log.Fatalf("con write err: %v", err)
		}
	}
}
