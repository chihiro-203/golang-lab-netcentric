package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printGrid(grid [][]rune) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Println()
	}
}

func generateGrid(rows, cols, mines int) [][]rune {
	rand.Seed(time.Now().UnixNano())

	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	placed := 0
	for placed < mines {
		r := rand.Intn(rows)
		c := rand.Intn(cols)
		if grid[r][c] != '*' {
			grid[r][c] = '*'
			placed++
		}
	}

	return grid
}

func aroundMines(grid [][]rune) [][]rune {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			count := 0
			if grid[i][j] == '.' {
				for _, dir := range directions {
					r, c := i+dir[0], j+dir[1]
					if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == '*' {
						count++
						if count > 0 {
							grid[i][j] = rune('0' + count)
						}
					}
				}
			}
		}
	}
	return grid
}

func main() {
	grid := generateGrid(20, 25, 99)
	printGrid(grid)
	fmt.Println()
	grid = aroundMines(grid)
	printGrid(grid)
}
