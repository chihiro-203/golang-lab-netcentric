package main

import (
	"bufio"
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
	}
}

func login(conn net.Conn) {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Logging in...")
		fmt.Print("Enter username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Enter password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		message := "/login " + username + " " + password
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

		if strings.Contains(strings.TrimSpace(response), "failed") {
			fmt.Printf("Server: %s", response)
			continue
		}

		response = strings.TrimSpace(response)

		key := strings.TrimSpace(response)

		fmt.Printf("%s_Login successfully. Your key: %s\n", key, response)
		fmt.Println("To continue, please follow instructions below:")
		fmt.Println("- Enter '/profile' to modify your profile")
		fmt.Println("- Enter '/game' to play Guessing Game")
		fmt.Println("- Enter '/file' to play download our file")

		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		message, _ = reader.ReadString('\n')
		message = strings.TrimSpace(message)

		writeMsg(conn, message)

		if message == "/profile" {
			modifyProfile(conn)
		} else if message == "/game" {
			guessingGame(conn)
		} else if message == "/file" {
			downloadFile(conn)
		}

		break
	}
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
			// hashed, err := hashPassword(password)
			// if err != nil {
			// 	fmt.Println("Error hashing password:", err)
			// }
			// message := "/register " + username + " " + hashed
			// _, err = conn.Write([]byte(message + "\n"))
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
			break
		} else {
			fmt.Println("Passwords do not match. Please try again.")
		}
	}

	hashed, err := hashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
	}
	message := "/register " + username + " " + string(hashed[:])
	_, err = conn.Write([]byte(message + "\n"))
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
}

func writeMsg(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("Error sending login data to server:", err)
		return
	}
}

func readMsg(conn net.Conn) string {
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return ""
	}
	return response
}

func modifyProfile(conn net.Conn) {
	fmt.Println("Modifying your profile...")

	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Username ('no' if not change): ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	writeMsg(conn, message)

	fmt.Println("Fullname ('no' if not change): ")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)
	writeMsg(conn, message)

	fmt.Println("Email ('no' if not change): ")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)
	writeMsg(conn, message)

	fmt.Println("Address ('no' if not change): ")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)
	writeMsg(conn, message)

	fmt.Println("Do you want to change password? \n'yes' if change\n'no' if not change")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)
	writeMsg(conn, message)
	if message == "yes" {
		fmt.Println("Old password: ")
		message, _ = reader.ReadString('\n')
		message = strings.TrimSpace(message)
		writeMsg(conn, message)

		fmt.Println("New password: ")
		message, _ = reader.ReadString('\n')
		message = strings.TrimSpace(message)
		writeMsg(conn, message)
	}
}

func guessingGame(conn net.Conn) {
	fmt.Println("Starting game...")

	var num string

gameLoop:
	for {
		fmt.Print("Input your guessed number: ")
		fmt.Scan(&num)

		writeMsg(conn, num)
		response := readMsg(conn)
		fmt.Print(response)
		os.Stdout.Sync()

		if strings.Contains(strings.ToLower(response), "correct") {
			answer := ""
			for answer != "yes" && answer != "no" {
				fmt.Println("Do you want to play again? (yes/no)")
				fmt.Scan(&answer)
				answer = strings.ToLower(answer)
				if answer == "yes" {
					continue
				} else if answer == "no" {
					fmt.Println("Thanks for playing the Guessing game.")
					break gameLoop
				} else {
					fmt.Println("Please type the correct syntax.")
				}
			}
			writeMsg(conn, answer)
		} else {
			continue
		}
	}

}

func downloadFile(conn net.Conn) {
	fmt.Println("Looking for available files...")
}
