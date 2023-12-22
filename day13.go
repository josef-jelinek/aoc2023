package main

import (
	"strings"
)

func solveDay13Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input)+"\n", "\n")
	var grid []string
	sum := 0
	for _, line := range lines {
		if line == "" {
			i := findSymLine(grid)
			if i > 0 {
				sum += 100 * i
			}
			sum += findSymLine(transpose(grid))
			grid = grid[:0]
			continue
		}
		grid = append(grid, line)
	}
	return sum
}

func solveDay13Part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input)+"\n", "\n")
	var grid []string
	sum := 0
	for _, line := range lines {
		if line == "" {
			i := findSmudgeSymLine(grid)
			if i > 0 {
				sum += 100 * i
			}
			sum += findSmudgeSymLine(transpose(grid))
			grid = grid[:0]
			continue
		}
		grid = append(grid, line)
	}
	return sum
}

func findSymLine(grid []string) int {
lineLoop:
	for i := 1; i < len(grid); i++ {
		n := min(i, len(grid)-i)
		for j := 0; j < n; j++ {
			if grid[i-j-1] != grid[i+j] {
				continue lineLoop
			}
		}
		return i
	}
	return 0
}

func findSmudgeSymLine(grid []string) int {
lineLoop:
	for i := 1; i < len(grid); i++ {
		n := min(i, len(grid)-i)
		smudges := 0
		for j := 0; j < n; j++ {
			smudges += count2Smudges(grid[i-j-1], grid[i+j])
			if smudges > 1 {
				continue lineLoop
			}
		}
		if smudges == 1 {
			return i
		}
	}
	return 0
}

func count2Smudges(s, t string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			n++
			if n > 1 {
				return n
			}
		}
	}
	return n
}

func transpose(grid []string) []string {
	t := make([]string, len(grid[0]))
	for _, s := range grid {
		for j := range s {
			t[j] += s[j : j+1]
		}
	}
	return t
}
