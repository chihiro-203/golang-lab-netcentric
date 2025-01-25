package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type User struct {
	Username          string   `json:"username"`
	EncryptedPassword string   `json:"encrypted"`
	FullName          string   `json:"fullname"`
	Email             []string `json:"email"`
	Address           []string `json:"address"`
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening to client:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}
		fmt.Println("Client connected!")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected!")
			return
		}

		fmt.Printf("Client chose to %s", message)

		response := "Server: "

		if message == "/login" {
			response += "Please login."
		} else if message == "/register" {
			response += "Please register."
		}
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}

		// response := "Server received: " + message
		// _, err = conn.Write([]byte(response))
		// if err != nil {
		// 	fmt.Println("Error writing to client:", err)
		// 	return
		// }
	}
}
