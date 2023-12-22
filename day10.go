package main

import (
	"strings"
)

func solveDay10Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	dimX := len(lines[0])
	dimY := len(lines)
	dists := make([][]int, dimY)
	var startX, startY int
	for y := 0; y < dimY; y++ {
		dists[y] = make([]int, dimX)
		for x := 0; x < dimX; x++ {
			dists[y][x] = -1
			if lines[y][x] == 'S' {
				startX, startY = x, y
			}
		}
	}
	nextX := []int{startX}
	nextY := []int{startY}
	nextD := []int{0}
	maxD := 0
	for len(nextD) > 0 {
		x, y, d := nextX[0], nextY[0], nextD[0]
		nextX, nextY, nextD = nextX[1:], nextY[1:], nextD[1:]
		if dists[y][x] >= 0 {
			continue
		}
		dists[y][x] = d
		maxD = max(maxD, d)
		// This is not correct as we do not consider connectivity, but works just for distances.
		if x > 0 && strings.IndexByte("-FL", lines[y][x-1]) >= 0 {
			nextX, nextY, nextD = append(nextX, x-1), append(nextY, y), append(nextD, d+1)
		}
		if x < dimX-1 && strings.IndexByte("-7J", lines[y][x+1]) >= 0 {
			nextX, nextY, nextD = append(nextX, x+1), append(nextY, y), append(nextD, d+1)
		}
		if y > 0 && strings.IndexByte("|F7", lines[y-1][x]) >= 0 {
			nextX, nextY, nextD = append(nextX, x), append(nextY, y-1), append(nextD, d+1)
		}
		if y < dimY-1 && strings.IndexByte("|LJ", lines[y+1][x]) >= 0 {
			nextX, nextY, nextD = append(nextX, x), append(nextY, y+1), append(nextD, d+1)
		}
	}
	return maxD
}

func solveDay10Part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	dimX := len(lines[0])
	dimY := len(lines)
	edges := make([][]byte, dimY)
	var startX, startY int
	for y := 0; y < dimY; y++ {
		edges[y] = make([]byte, dimX)
		for x := 0; x < dimX; x++ {
			if lines[y][x] == 'S' {
				startX, startY = x, y
			}
		}
	}
	// We need to replace S by the right pipe, so our inside/outside tracking works.
	tcon := startY > 0 && strings.IndexByte("|F7", lines[startY-1][startX]) >= 0
	bcon := startY < dimY-1 && strings.IndexByte("|LJ", lines[startY+1][startX]) >= 0
	lcon := startX > 0 && strings.IndexByte("-FL", lines[startY][startX-1]) >= 0
	rcon := startX < dimX-1 && strings.IndexByte("-7J", lines[startY][startX+1]) >= 0
	var start string
	switch {
	case lcon && rcon:
		start = "-"
	case tcon && bcon:
		start = "|"
	case rcon && bcon:
		start = "F"
	case lcon && bcon:
		start = "7"
	case rcon && tcon:
		start = "L"
	case lcon && tcon:
		start = "J"
	}
	lines[startY] = strings.Replace(lines[startY], "S", start, 1)
	nextX := []int{startX}
	nextY := []int{startY}
	for len(nextX) > 0 {
		x, y := nextX[0], nextY[0]
		nextX, nextY = nextX[1:], nextY[1:]
		if edges[y][x] > 0 {
			continue
		}
		edges[y][x] = 1
		c := lines[y][x]
		// We care about true loop edges and need to check for connectivity unline in part 1.
		if x > 0 && strings.IndexByte("-7J", c) >= 0 && strings.IndexByte("-FL", lines[y][x-1]) >= 0 {
			nextX, nextY = append(nextX, x-1), append(nextY, y)
		}
		if x < dimX-1 && strings.IndexByte("-FL", c) >= 0 && strings.IndexByte("-7J", lines[y][x+1]) >= 0 {
			nextX, nextY = append(nextX, x+1), append(nextY, y)
		}
		if y > 0 && strings.IndexByte("|LJ", c) >= 0 && strings.IndexByte("|F7", lines[y-1][x]) >= 0 {
			nextX, nextY = append(nextX, x), append(nextY, y-1)
		}
		if y < dimY-1 && strings.IndexByte("|F7", c) >= 0 && strings.IndexByte("|LJ", lines[y+1][x]) >= 0 {
			nextX, nextY = append(nextX, x), append(nextY, y+1)
		}
	}
	count := 0
	for y := 0; y < dimY; y++ {
		line := lines[y]
		lineEdges := edges[y]
		inside := false
		for x := 0; x < dimX; x++ {
			if lineEdges[x] == 0 {
				if inside {
					count++
				}
				continue
			}
			switch line[x] {
			case '|':
				inside = !inside
			case 'F':
				x++
				for line[x] == '-' {
					x++
				}
				if line[x] == 'J' {
					inside = !inside
				}
			case 'L':
				x++
				for line[x] == '-' {
					x++
				}
				if line[x] == '7' {
					inside = !inside
				}
			}
		}
	}
	return count
}
