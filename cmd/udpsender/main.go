package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func main(){
	address, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil{
		fmt.Println("Unable to resolve address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, address)
	if err != nil{
		fmt.Println("Unable to establish connection:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Unable to read string:", err)
			return
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Unable to send data:", err)
			return
		}
	}
}