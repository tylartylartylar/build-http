package main

import (
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
	for {
		_, err := file.Read(data)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		fmt.Printf("read: %s\n", data)
	}

}
