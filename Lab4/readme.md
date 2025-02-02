### [Return to README.md](https://github.com/ume-meu/golang-lab-netcentric/blob/main/README.md)

# Lab 4 - TCP Socket

This repository contains the lab exercises for Net-Centric Programming, focusing on TCP socket programming using Golang. Below are the objectives and exercises included in this lab.

## Objective
- Learn the basics of TCP socket programming in Go.

---

## Exercises

- Server maintain a list of user record, each user contains the following information: username, password, full name,
list of email, list of address (home address, working address,…)
- Password must be encrypted, Base64 (or any) can be used
- All user record must be save in .JSON or .GOB file
- Authentication is needed when client connect to the server
- Server also generate a random integer value and send back to client after authentication is finished. This key is unique among clients
- Message exchange between client server need to attach this key as a prefix. Example: 125_Hello server , 125_Hello client

Simple hangman game (similar to “Chiec non ky dieu” game show on VTV)
- Server main tain a list of word and its description. Can be stored in file
- Unrevealed letter must be display using underscore _ symbol
- 2 or more players can play this game
- Each player has 30 second to guess a letter
- If the guessed letter is correct. Player will get 10pts x number of appearance of the letter. Player
can continue the next guess
- If the guessed is wrong, player loose his/her turn and the other player can make the guess
- The game ends when only player finish revealing the hidden word

---

## Notes
These exercises are designed to help you get comfortable with TCP socket programming in Go, including user authentication and multiplayer interaction. Feel free to contribute improvements or share your solutions!
