package main

import (
	"fmt"
	"strings"
)

func solveDay14Part1(input string) {
	grid := byteGrid(strings.Split(strings.TrimSpace(input), "\n"))
	sum := 0
	for col := 0; col < len(grid[0]); col++ {
		pos := 0
		for row := 0; row < len(grid); row++ {
			switch grid[row][col] {
			case '#':
				pos = row + 1
			case 'O':
				sum += len(grid) - pos
				pos++
			}
		}
	}
	fmt.Printf("Day 14, Problem 1, Answer: %v\n", sum)
}

func solveDay14Part2(input string) {
	const numCycles = 1_000_000_000
	grid := byteGrid(strings.Split(strings.TrimSpace(input), "\n"))
	w, h := len(grid[0]), len(grid)
	var hashHistory []int
	var cycleStart, cycleEnd, cycleIndex int
	load := 0
	for cycle := 0; cycle < numCycles; cycle++ {
		doCycle(grid, w, h)
		hash := getHash(grid, w, h)
		if cycleIndex > 0 {
			if hashHistory[cycleStart+cycleIndex] == hash {
				if cycleIndex == cycleEnd-1 {
					cycleLen := cycleEnd - cycleStart + 1
					todo := (numCycles - 1 - cycle) % cycleLen
					for i := 0; i < todo; i++ {
						doCycle(grid, w, h)
					}
					load = northBeamLoad(grid, w, h)
					break
				}
				cycleIndex++
			} else {
				cycleIndex = 0
			}
		}
		if cycleIndex == 0 {
			for i := len(hashHistory) - 1; i >= 0; i-- {
				if hashHistory[i] == hash {
					cycleStart = i
					cycleEnd = len(hashHistory) - 1
					cycleIndex = 1
				}
			}
		}
		hashHistory = append(hashHistory, hash)
	}
	fmt.Printf("Day 14, Problem 2, Answer: %v\n", load)
}

func doCycle(grid [][]byte, w, h int) {
	for x := 0; x < w; x++ {
		i := 0
		for y := 0; y < h; y++ {
			switch grid[y][x] {
			case '#':
				i = y + 1
			case 'O':
				grid[y][x] = '.'
				grid[i][x] = 'O'
				i++
			}
		}
	}
	for y := 0; y < h; y++ {
		i := 0
		for x := 0; x < w; x++ {
			switch grid[y][x] {
			case '#':
				i = x + 1
			case 'O':
				grid[y][x] = '.'
				grid[y][i] = 'O'
				i++
			}
		}
	}
	for x := 0; x < w; x++ {
		i := h - 1
		for y := h - 1; y >= 0; y-- {
			switch grid[y][x] {
			case '#':
				i = y - 1
			case 'O':
				grid[y][x] = '.'
				grid[i][x] = 'O'
				i--
			}
		}
	}
	for y := 0; y < h; y++ {
		i := w - 1
		for x := w - 1; x >= 0; x-- {
			switch grid[y][x] {
			case '#':
				i = x - 1
			case 'O':
				grid[y][x] = '.'
				grid[y][i] = 'O'
				i--
			}
		}
	}
}

func getHash(grid [][]byte, w, h int) int {
	hash := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 'O' {
				hash += y*w + x
			}
		}
	}
	return hash
}

func northBeamLoad(grid [][]byte, w, h int) int {
	sum := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 'O' {
				sum += h - y
			}
		}
	}
	return sum
}

func byteGrid(lines []string) [][]byte {
	t := make([][]byte, len(lines))
	for i, s := range lines {
		t[i] = []byte(s)
	}
	return t
}
