package main

import (
	"io"
)

type Terrain byte

const (
	Land        = Terrain('1')
	OutOfBounds = Terrain(255)
	Visited     = Terrain(254)
)

type Grid [][]Terrain

func (grid Grid) At(x, y int) Terrain {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return OutOfBounds
	}
	return grid[y][x]
}

func (grid Grid) Set(x, y int, value Terrain) {
	grid[y][x] = value
}

func FromReader(source io.Reader) (Grid, error) {
	filebuff := make([]byte, 8192)
	grid := Grid{}
	row := []Terrain{}

	n, err := source.Read(filebuff)
	for err == nil {
		for i := 0; i < n; i++ {
			if filebuff[i] == '\n' {
				grid = append(grid, row)
				row = make([]Terrain, 0, len(row))
				continue
			}
			row = append(row, Terrain(filebuff[i]))
		}
		n, err = source.Read(filebuff)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}

	for i := 0; i < n; i++ {
		if filebuff[i] == '\n' {
			grid = append(grid, row)
			row = make([]Terrain, 0, len(row))
			continue
		}
		row = append(row, Terrain(filebuff[i]))
	}
	if len(row) > 0 {
		grid = append(grid, row)
	}
	return grid, nil
}
