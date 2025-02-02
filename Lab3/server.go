package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	authenticated     = make(map[string]int)
	isClientConnected bool
	mu                sync.Mutex
	userFile          = "users.json"
	name              string
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

	mu.Lock()
	isClientConnected = true
	defer func() {
		isClientConnected = false
		mu.Unlock()
	}()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected!")
			return
		}
		message = strings.TrimSpace(message)
		parts := strings.Split(message, " ")
		action, username, password := parts[0], parts[1], parts[2]

		if action == "/register" {
			if userRegister(username, password) {
				conn.Write([]byte("Registration successful. Please login.\n"))
			} else {
				conn.Write([]byte("Registration failed. Username may already exist.\n"))
			}
		} else if action == "/login" {
			prefix := userLogin(username, password)
			if prefix != 0 {
				fmt.Printf("Client's username: %s\n", username)
				name = username
				authenticated[username] = prefix
				conn.Write([]byte(fmt.Sprintf("%d\n", prefix)))
				message = receiveMsg(conn)
				if message == "/profile" {
					modifyInfo()
				} else if message == "/game" {
					playGame(conn)
				} else if message == "/file" {
					getFiles(conn)
				}
			} else {
				conn.Write([]byte("Login failed. Invalid username or password.\n"))
			}
		} else {
			conn.Write([]byte("Unknown command.\n"))
		}
	}
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

	prefix := genNum()

	newUser := User{
		Username: username,
		Password: password,
		Prefix:   prefix,
	}

	users = append(users, newUser)
	saveUser(users)

	return true
}

func userLogin(username, password string) int {
	users := loadUsers()

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	for _, user := range users {
		if user.Username == username && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil {
			return user.Prefix
		}
	}

	return 0
}

func modifyInfo() {}

func playGame(conn net.Conn) {
	randNum := genNum()
	fmt.Printf("The random number is: %d\n", randNum)

	for {
		reader := bufio.NewReader(conn)
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		answer, _ := strconv.Atoi(message)

		if answer > randNum {
			conn.Write([]byte("Your guessed number is larger than the result. Please try again.\n"))
		} else if answer < randNum {
			conn.Write([]byte("Your guessed number is smaller than the result. Please try again.\n"))
		} else {
			conn.Write([]byte(fmt.Sprintf("Correct answer. The random number is %d.\n", randNum)))

			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(response)
			os.Stdout.Sync()
			if response == "yes" {
				randNum = genNum()
				fmt.Printf("The random number is: %d\n", randNum)
				continue
			} else if response == "no" {
				fmt.Println(name, "finishes playing game.")
				break
			}
		}
	}
}

func getFiles(conn net.Conn) {
	folderPath := "./files"

	for {
		files, err := os.ReadDir(folderPath)
		if err != nil {
			fmt.Println("Error reading folder:", err)
		}

		var names []string

		for _, file := range files {
			if !file.IsDir() {
				names = append(names, file.Name())
			}
		}

		namesString := strings.Join(names, "  ")
		conn.Write([]byte(namesString + "\n"))

		filename := receiveMsg(conn)
		fmt.Println(authenticated[name], "_I choose", filename)

		filename = filepath.Join(folderPath, filename)
		file, err := os.Open(filename)
		if err != nil {
			conn.Write([]byte("File not found.\n"))
			fmt.Println("File not found:", filename)
			return
		}
		defer file.Close()

		_, err = io.Copy(conn, file)
		if err != nil {
			fmt.Println("Error sending file:", err)
			return
		}

		fmt.Println("File sent successfully:", filename)
		conn.Write([]byte("EOF\n"))

		conn.Write([]byte("File transfer complete.\n"))

		answer := receiveMsg(conn)
		if answer != "yes" {
			fmt.Println("Client chose not to download more files.")
			break
		}
	}

}

func genNum() int {
	// rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

func receiveMsg(conn net.Conn) string {
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return ""
	}
	response = strings.TrimSpace(response)
	return response
}
