package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := make([]byte, 8)
	currentLine := ""
	reader := bufio.NewReader(file)

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
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
			} else {
				currentLine += string(part)
			}
		}
	}

	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
