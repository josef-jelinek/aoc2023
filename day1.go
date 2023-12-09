package main

import (
	"fmt"
	"strings"
)

func solveDay1Part1(input string) {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		var d10, d1 int
		d10Set := false
		for _, r := range s {
			if '0' <= r && r <= '9' {
				d1 = int(r - '0')
				if !d10Set {
					d10 = 10 * d1
					d10Set = true
				}
			}
		}
		sum += d10 + d1
	}
	fmt.Printf("Day 1, Problem 1, Answer: %d\n", sum)
}

func solveDay1Part2(input string) {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		var d10, d1 int
		d10Set := false
		for i, r := range s {
			d := 0
			switch {
			case '0' <= r && r <= '9':
				d = int(r - '0')
			case strings.HasPrefix(s[i:], "one"):
				d = 1
			case strings.HasPrefix(s[i:], "two"):
				d = 2
			case strings.HasPrefix(s[i:], "three"):
				d = 3
			case strings.HasPrefix(s[i:], "four"):
				d = 4
			case strings.HasPrefix(s[i:], "five"):
				d = 5
			case strings.HasPrefix(s[i:], "six"):
				d = 6
			case strings.HasPrefix(s[i:], "seven"):
				d = 7
			case strings.HasPrefix(s[i:], "eight"):
				d = 8
			case strings.HasPrefix(s[i:], "nine"):
				d = 9
			default:
				continue
			}
			d1 = d
			if !d10Set {
				d10 = 10 * d
				d10Set = true
			}
		}
		sum += d10 + d1
	}
	fmt.Printf("Day 1, Problem 2, Answer: %d\n", sum)
}
