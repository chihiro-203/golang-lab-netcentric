package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	authenticated     = make(map[string]bool)
	isClientConnected bool
	mu                sync.Mutex
	userFile          = "users.json"
)

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	FullName string   `json:"fullname"`
	Email    []string `json:"email"`
	Address  []string `json:"address"`
	Prefix   int      `json:"prefix"`
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

		// Check if any clients try to connect.
		mu.Lock()
		if isClientConnected {
			fmt.Println("New client tried to connect, but a client is already connected.")
			conn.Write([]byte("Server is busy. Try again later.\n"))
			conn.Close()
			mu.Unlock()
			continue
		}
		isClientConnected = true
		mu.Unlock()

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
		parts := strings.Split(message, " ")
		action, username, password := parts[0], parts[1], parts[2]

		user := User{
			Username: username,
			Password: password,
		}

		if action == "/login" {
		} else if action == "/register" {
		}
		fmt.Printf("Client joins: %s", username)
	}

	mu.Lock()
	isClientConnected = false
	mu.Unlock()
}

func loadUsers() []User {
	var users []User

	if _, err := os.Stat(userFile); os.IsNotExist(err) {
		return users
	}

	data, err := ioutil.ReadFile(userFile)
	if err != nil {
		fmt.Println("Error reading user file:", err)
		return users
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error decoding user JSON:", err)
		return users
	}

	return users
}

func saveUser(users []User) {
	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = ioutil.WriteFile(userFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing to user file:", err)
	}
}

func userRegister(username, password string) bool {
	users := loadUsers()

	for _, user := range users {
		if user.Username == username {
			return false
		}
	}

	prefix := genPrefix()

	newUser := User{
		Username: username,
		Password: password,
		Prefix:   prefix,
	}
	return true
}

func genPrefix() int {
	// rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}
