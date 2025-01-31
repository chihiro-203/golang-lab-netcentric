package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

var hash = sha256.New()

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server.")
	fmt.Println("You need to login or sign up to send message.")
	fmt.Println("- Enter '/login' to log in")
	fmt.Println("- Enter '/register' to sign up")

	// authenticated := false

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "/login" {
			login(conn)
			break
		} else if message == "/register" {
			register(conn)
			break
		} else {
			fmt.Println("You need to login or sign up to send a message.")
			fmt.Println("- Enter '/login' to log in")
			fmt.Println("- Enter '/register' to sign up")
		}

		// if message != "/login" && message != "/register" {
		// 	fmt.Println("You need to login or sign up to send message.")
		// 	fmt.Println("- Enter '/login' to log in")
		// 	fmt.Println("- Enter '/register' to sign up")
		// } else {

		// _, err := conn.Write([]byte(message))
		// if err != nil {
		// 	fmt.Println("Error writing to server:", err)
		// 	os.Exit(1)
		// }

		// response, err := bufio.NewReader(conn).ReadString('\n')
		// if err != nil {
		// 	fmt.Println("Error reading from server:", err)
		// 	return
		// }

		// fmt.Printf("Server response %s", response)
		// }
	}
}

func login(conn net.Conn) {
	fmt.Println("Logging in...")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	hash.Write([]byte(password))
	hashed := hash.Sum(nil)
	message := "/login " + username + " " + string(hashed[:])
	_, err := conn.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("Error sending login data to server:", err)
		return
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Printf("Server: %s", response)
}

func register(conn net.Conn) {
	fmt.Println("Registering a new account...")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	var password, cPassword string
	for {
		fmt.Print("Enter password: ")
		password, _ = reader.ReadString('\n')
		password = strings.TrimSpace(password)

		fmt.Print("Confirm password: ")
		cPassword, _ = reader.ReadString('\n')
		cPassword = strings.TrimSpace(cPassword)

		if password == cPassword {
			hash.Write([]byte(password))
			hashed := hash.Sum(nil)
			message := "/register " + username + " " + string(hashed[:])
			_, err := conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error sending registration data to server:", err)
				return
			}

			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from server:", err)
				return
			}
			fmt.Printf("Server:  %s", response)
			break
		} else {
			fmt.Println("Passwords do not match. Please try again.")
		}
	}

	// hash.Write([]byte(password))
	// hashed := hash.Sum(nil)
	// message := "/register " + username + " " + string(hashed[:])
	// _, err := conn.Write([]byte(message + "\n"))
	// if err != nil {
	// 	fmt.Println("Error sending registration data to server:", err)
	// 	return
	// }

	// response, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading from server:", err)
	// 	return
	// }
	// fmt.Printf("Server:  %s", response)
}

func modifyInfo(conn net.Conn) {}

func guessingGame(conn net.Conn) {}

func downloadFile(conn net.Conn) {}
