package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filename = "messages.txt"

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
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s: %v", filename, err)
	}
	defer file.Close()

	c := getLinesChannel(file)
	for item := range c {
		fmt.Printf("read: %s\n", item)
	}
}
