package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveDay9Part1(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, s := range lines {
		if s == "" {
			continue
		}
		numStrs := strings.Split(s, " ")
		numRows := [][]int{make([]int, len(numStrs))}
		for i, numStr := range numStrs {
			numRows[0][i], _ = strconv.Atoi(numStr)
		}
		done := false
		i := 0
		for !done {
			done = true
			numRows = append(numRows, make([]int, len(numRows[i])-1))
			for j := 0; j < len(numRows[i])-1; j++ {
				d := numRows[i][j+1] - numRows[i][j]
				numRows[i+1][j] = d
				if d != 0 {
					done = false
				}
			}
			i++
		}
		next := 0
		for i >= 0 {
			next += numRows[i][len(numRows[i])-1]
			i--
		}
		sum += next
	}
	fmt.Printf("Day 9, Problem 1, Answer: %v\n", sum)
}

func solveDay9Part2(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, s := range lines {
		if s == "" {
			continue
		}
		numStrs := strings.Split(s, " ")
		numRows := [][]int{make([]int, len(numStrs))}
		for i, numStr := range numStrs {
			numRows[0][i], _ = strconv.Atoi(numStr)
		}
		done := false
		i := 0
		for !done {
			done = true
			numRows = append(numRows, make([]int, len(numRows[i])-1))
			for j := 0; j < len(numRows[i])-1; j++ {
				d := numRows[i][j+1] - numRows[i][j]
				numRows[i+1][j] = d
				if d != 0 {
					done = false
				}
			}
			i++
		}
		prev := 0
		for i >= 0 {
			prev = numRows[i][0] - prev
			i--
		}
		sum += prev
	}
	fmt.Printf("Day 9, Problem 2, Answer: %v\n", sum)
}
