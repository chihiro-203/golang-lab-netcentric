### [Return to README.md](https://github.com/ume-meu/golang-lab-netcentric/blob/main/README.md)

# Lab 1 - Introduction to Golang

This repository contains the solutions to the lab exercises for Net-Centric Programming, focusing on an introduction to the Go programming language. Below are the objectives and exercises included in this lab.

## Objective
- Familiarize yourself with the Go Programming Language.

---

## Exercises

### 1. Hamming Distance ([Solution](https://github.com/ume-meu/golang-lab-netcentric/blob/main/Lab1/Ex1/main.go))

**Problem:**
Calculate the Hamming Distance between two DNA strands.

**Details:**
DNA strands are represented using the letters C, A, G, and T. The Hamming Distance is calculated by comparing two DNA strands of equal length and counting the differences between them.

- **Example**:
  ```
  DNA Strand 1: GAGCCTACTAACGGGAT
  DNA Strand 2: CATCGTAATGACGGCCT
                 ^ ^ ^^ ^    ^^
  Hamming Distance: 7
  ```

- **Note:** The Hamming Distance is undefined for DNA strands of unequal length.
- **Task:** Generate 1,000 pairs of random DNA samples of a given length and compute their Hamming Distance.

---

### 2. Scrabble Score ([Solution](https://github.com/ume-meu/golang-lab-netcentric/blob/main/Lab1/Ex2/main.go))
**Problem:**
Given a word, compute its Scrabble score based on the following letter values:

| Letters                   | Value |
|---------------------------|-------|
| A, E, I, O, U, L, N, R, S, T | 1|
| D, G                     | 2     |
| B, C, M, P               | 3     |
| F, H, V, W, Y            | 4     |
| K                        | 5     |
| J, X                     | 8     |
| Q, Z                     | 10    |

- **Example:**
  ```
  Word: "cabbage"
  Scrabble Score: 14
  Explanation:
  - C: 3 points
  - A: 1 point (twice)
  - B: 3 points (twice)
  - G: 2 points
  - E: 1 point
  Total: 3 + 2*1 + 2*3 + 2 + 1 = 14
  ```
- Extend the task to compute the score for a string containing multiple words.

---

### 3. Luhn Algorithm ([Solution](https://github.com/ume-meu/golang-lab-netcentric/blob/main/Lab1/Ex3/main.go))
**Problem:**
Validate whether a given number satisfies the Luhn formula, commonly used for credit card validation.

**Rules:**
- Ignore spaces and validate only numeric strings.
- Strings of length 1 or less are invalid.
- Double every second digit from the right. If the result exceeds 9, subtract 9.
- Sum all the digits. The number is valid if the sum is divisible by 10.

- **Example (Valid Number):**  
  ```
  Input: 4539 3195 0343 6467
  Process: Double and subtract if necessary:
           8 5 6 9 6 1 9 5 0 3 8 3 3 4 3 7
  Sum: 80 (Divisible by 10)
  Valid
  ```

- **Example (Invalid Number):**  
  ```
  Input: 8273 1232 7352 0569
  Process: Double and subtract if necessary:
           7 2 5 3 2 2 6 2 5 3 1 2 0 5 3 9
  Sum: 57 (Not divisible by 10)
  Invalid
  ```

---

### 4. Minesweeper ([Solution](https://github.com/ume-meu/golang-lab-netcentric/blob/main/Lab1/Ex4/main.go))
**Problem:**  
Add mine counts to a completed Minesweeper board.

**Details:**  
- Input: A minefield where `*` represents a mine and `.` represents an empty square.
- Output: Replace each empty square with the count of adjacent mines. Leave empty squares with no adjacent mines blank.

- **Example (5x4 Board):**
  ```
  Input:
  .*.*.
  ..*..
  ..*..
  .....

  Output:
  1*3*1
  13*31
  .2*2.
  .111.
  ```
- **Task:** Create a random 20x25 minefield with 99 mines and mark the board.

---

### 5. Matching Brackets ([Solution](https://github.com/ume-meu/golang-lab-netcentric/blob/main/Lab1/Ex5/main.go))
**Problem:**
Given a string containing brackets `[ ]`, braces `{ }`, and parentheses `( )`, verify that all pairs are matched and nested correctly.

- **Example:**
  ```
  Input: fmt.Println(a.TypeOf(xyz)){[ ]}
  Output: Correct
  ```
