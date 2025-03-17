package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()

		data := make([]byte, 8)
		currentLine := ""
		reader := bufio.NewReader(f)

		for {
			n, err := reader.Read(data)
			if err != nil {
				if err.Error() != "EOF" {
					fmt.Println("Error reading file:", err)
				}
				break
			}

			parts := string(data[:n])
			for _, part := range parts {
				if part == '\n' {
					lines <- currentLine
					currentLine = ""
				} else {
					currentLine += string(part)
				}
			}
		}

		if currentLine != "" {
			lines <- currentLine
		}
		close(lines)
	}()
	return lines
}

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	linesChannel := getLinesChannel(file)
	for line := range linesChannel {
		fmt.Printf("read: %s\n", line)
	}
}
