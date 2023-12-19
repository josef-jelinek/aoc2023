package main

import (
	"fmt"
	"slices"
	"strings"
)

func solveDay4Part1(input string) {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		_, numsPart, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		wantPart, gotPart, _ := strings.Cut(numsPart, " | ")
		wantNums := strings.Split(wantPart, " ")
		gotNums := strings.Split(gotPart, " ")
		count := 0
		for _, gotNum := range gotNums {
			if gotNum != "" && slices.Contains(wantNums, gotNum) {
				count *= 2
				if count == 0 {
					count = 1
				}
			}
		}
		sum += count
	}
	fmt.Printf("Day 4, Problem 1, Answer: %d\n", sum)
}

func solveDay4Part2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	won := make([]int, len(lines))
	for i, s := range lines {
		_, numsStr, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		wantNumsStr, gotNumsStr, _ := strings.Cut(numsStr, " | ")
		wantNums := strings.Split(wantNumsStr, " ")
		gotNums := strings.Split(gotNumsStr, " ")
		count := 0
		for _, gotNum := range gotNums {
			if gotNum != "" && slices.Contains(wantNums, gotNum) {
				count++
			}
		}
		for j := 0; j < count; j++ {
			won[i+j+1] += 1 + won[i]
		}
	}
	sum := 0
	for _, w := range won {
		sum += 1 + w
	}
	fmt.Printf("Day 4, Problem 2, Answer: %d\n", sum)
}
