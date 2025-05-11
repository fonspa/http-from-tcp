package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const address = ":42069"

func getLinesChannel(f io.ReadCloser) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		line := ""
		for {
			b := make([]byte, 8)
			n, err := f.Read(b)
			if err != nil {
				if errors.Is(err, io.EOF) {
					if line != "" {
						c <- line
					}
					return
				}
				log.Fatalf("unable to read bytes from file: %v", err)
			}
			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := range len(parts) - 1 {
				c <- fmt.Sprintf("%s%s", line, parts[i])
				line = ""
			}
			line += parts[len(parts)-1]
		}
	}()
	return c
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("unable to setup tcp listener: %v", err)
	}
	defer listener.Close()
	con, err := listener.Accept()
	if err != nil {
		log.Fatalf("unable to connect to the TCP listener: %v", err)
	}
	defer con.Close()
	fmt.Println("Connection accepted")

	c := getLinesChannel(con)
	for item := range c {
		fmt.Printf("%s\n", item)
	}
	fmt.Println("Connection closed")
}
