package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Enter message: ")

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			os.Exit(1)
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		fmt.Printf("Server response %s", response)
	}
}
