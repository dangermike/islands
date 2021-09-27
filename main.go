package main

import (
	"fmt"
	"os"
)

type Coord struct {
	X int
	Y int
}

func (c Coord) TransC(other Coord) Coord {
	return Coord{c.X + other.X, c.Y + other.Y}
}

var adjacencies = []Coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func visit(g Grid, x int, y int) int {
	count := 0
	if g.At(x, y) != Land {
		return 0
	}
	stack := []Coord{{x, y}}

	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if g.At(c.X, c.Y) == Land {
			count++
			g.Set(c.X, c.Y, Visited)
			for _, t := range adjacencies {
				stack = append(stack, c.TransC(t))
			}
		}
	}

	return count
}

func countIslands(grid Grid) int {
	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if 0 < visit(grid, x, y) {
				count++
			}
		}
	}
	return count
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(os.Args[i])
		fmt.Print(": ")
		f, err := os.Open(os.Args[i])
		if err != nil {
			fmt.Printf("Failed to open file: %v\n", err)
			os.Exit(1)
		}
		grid, err := FromReader(f)
		if err != nil {
			fmt.Printf("Failed to read file: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(countIslands(grid))
	}
	if len(os.Args) == 1 {
		grid, err := FromReader(os.Stdin)
		if err != nil {
			fmt.Printf("Failed to read stdin: %v\n", err)
		}
		if len(grid) > 0 {
			fmt.Println(countIslands(grid))
		}
	}
}
