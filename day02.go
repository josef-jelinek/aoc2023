package main

import (
	"strconv"
	"strings"
)

func solveDay2Part1(input string) any {
	maxByColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0
gameLoop:
	for _, s := range strings.Split(input, "\n") {
		name, spec, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		id, _ := strconv.Atoi(name[5:])
		for _, draw := range strings.Split(spec, "; ") {
			for _, cube := range strings.Split(draw, ", ") {
				count, color, _ := strings.Cut(cube, " ")
				n, _ := strconv.Atoi(count)
				if n > maxByColor[color] {
					continue gameLoop
				}
			}
		}
		sum += id
	}
	return sum
}

func solveDay2Part2(input string) any {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		_, spec, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		minByColor := make(map[string]int)
		for _, draw := range strings.Split(spec, "; ") {
			for _, cube := range strings.Split(draw, ", ") {
				count, color, _ := strings.Cut(cube, " ")
				n, _ := strconv.Atoi(count)
				minByColor[color] = max(minByColor[color], n)
			}
		}
		power := 1
		for _, v := range minByColor {
			power *= v
		}
		sum += power
	}
	return sum
}
