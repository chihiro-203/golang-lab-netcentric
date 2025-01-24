### [Return to README.md](https://github.com/ume-meu/golang-lab-netcentric/blob/main/README.md)

# Lab 3 - Simple TCP Socket

This repository contains the lab exercises for Net-Centric Programming, focusing on simple TCP socket programming using Golang. Below are the objectives and exercises included in this lab.

## Objective
- Learn and get familiar with TCP Socket Programming using Golang.

---

## Exercises

- Server maintain a list of user record, each user contains the following information: username, password, full name,
list of email, list of address (home address, working address,…)
- Password must be encrypted, Base64 (or any) can be used
- All user record must be save in .JSON or .GOB file
- Authentication is needed when client connect to the server
- Server also generate a random integer value and send back to client after authentication is finished. This key is
unique among clients
- Message exchange between client server need to attach this key as a prefix. Example: 125_Hello server , 125_Hello
client

### 1. Guessing game:
- As the game start, server generates a random number between 1 and 100
- Client sends a guessed number input by user
- Server sends back the indication where as the guessed number is larger or smaller than the result
- This progress stop when user input the exact number
- The game can be repeated or ended

### 2. Improve the game above to support authentication and prefix data

### 3. File download:
- Client send request for the file that need downloaded (i.e. file name)
- Server read the file content and send back to client
- Only simple text file is supported. The file size can be large, upto 10MB
- Optional: The program support user authentication and prefix

---

## Notes
These exercises are designed to help you get comfortable with TCP socket programming in Go and to understand key concepts like authentication, message prefixing, and file handling. Feel free to contribute improvements or share your solutions!