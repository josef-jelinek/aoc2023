package main

import (
	"fmt"
	"strings"
)

func solveDay21Part1(input string) {
	n := 64
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var sx, sy int
	for y := range lines {
		if x := strings.IndexByte(lines[y], 'S'); x >= 0 {
			sx, sy = x, y
			break
		}
	}
	plots := map[[2]int]bool{{sx, sy}: true}
	for i := 0; i < n; i++ {
		newPlots := make(map[[2]int]bool)
		for xy := range plots {
			x, y := xy[0], xy[1]
			if lines[y][x-1] != '#' {
				newPlots[[2]int{x - 1, y}] = true
			}
			if lines[y-1][x] != '#' {
				newPlots[[2]int{x, y - 1}] = true
			}
			if lines[y][x+1] != '#' {
				newPlots[[2]int{x + 1, y}] = true
			}
			if lines[y+1][x] != '#' {
				newPlots[[2]int{x, y + 1}] = true
			}
		}
		plots = newPlots
	}
	fmt.Printf("Day 21, Problem 1, Answer: %v\n", len(plots))
}

func solveDay21Part2(input string) {
	steps := 26501365
	dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	n := len(lines)
	var sx, sy int
	for y := range lines {
		if n != len(lines[y]) {
			panic("Unsupported input")
		}
		if x := strings.IndexByte(lines[y], 'S'); x >= 0 {
			sx, sy = x, y
		}
	}
	plots := map[[2]int]bool{{sx, sy}: true}
	var counts []int
	for i := 1; len(counts) < 3; i++ {
		newPlots := make(map[[2]int]bool)
		for xy := range plots {
			x, y := xy[0], xy[1]
			for _, dxy := range dirs {
				dx, dy := dxy[0], dxy[1]
				nx, ny := ((x+dx)%n+n)%n, ((y+dy)%n+n)%n // remaider -> modulo
				if lines[ny][nx] != '#' {
					newPlots[[2]int{x + dx, y + dy}] = true
				}
			}
		}
		plots = newPlots
		if i%n == steps%n {
			counts = append(counts, len(plots))
		}
	}
	// Constant offset added to the value.
	c := counts[0]
	// The first "derivative" increasing the value by this each step.
	d1 := counts[1] - c
	// The second "derivative" increasing the value by this times the step every 2 steps.
	// Contribution starts one step later than d1.
	d2 := counts[2] - counts[1] - d1 // increase by this + previous every two steps
	// 0 -> c
	// 1 -> 1*(0*d2/2 + d1) + c
	// 2 -> 2*(1*d2/2 + d1) + c
	// 3 -> 3*(2*d2/2 + d1) + c
	// 4 -> 4*(3*d2/2 + d1) + c
	// 5 -> 5*(4*d2/2 + d1) + c
	// 6 -> 6*(5*d2/2 + d1) + c
	// 7 -> 7*(6*d2/2 + d1) + c
	// ...
	x := steps / n
	y := x*((x-1)*d2/2+d1) + c
	fmt.Printf("Day 21, Problem 2, Answer: %v\n", y)
}
