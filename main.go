package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "messages.txt"

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s: %v", filename, err)
	}
	defer file.Close()

	for {
		b := make([]byte, 8)
		n, err := file.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalf("unable to read bytes from file: %v", err)
		}
		fmt.Printf("read: %s\n", string(b[:n]))
	}
}
