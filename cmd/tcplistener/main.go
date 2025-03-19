package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Connection accepted")

		linesChannel := getLinesChannel(conn)

		for line := range linesChannel {
			fmt.Println(line)
		}

		fmt.Println("Connection closed")
	}
}
