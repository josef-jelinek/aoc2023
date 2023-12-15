package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveDay15Part1(input string) {
	steps := strings.Split(strings.TrimSpace(input), ",")
	sum := 0
	for _, s := range steps {
		var h byte
		for _, b := range []byte(s) {
			h = (h + b) * 17
		}
		sum += int(h)
	}
	fmt.Printf("Day 15, Problem 1, Answer: %v\n", sum)
}

func solveDay15Part2(input string) {
	type lens struct {
		label string
		value int
	}
	steps := strings.Split(strings.TrimSpace(input), ",")
	boxes := make([][]lens, 256)
stepLoop:
	for _, s := range steps {
		var h byte
		label, valueStr, add := strings.Cut(s, "=")
		if !add {
			label = strings.TrimSuffix(s, "-")
		}
		for _, b := range []byte(label) {
			h = (h + b) * 17
		}
		lenses := boxes[h]
		if add {
			value, _ := strconv.Atoi(valueStr)
			for i := 0; i < len(lenses); i++ {
				if lenses[i].label == label {
					lenses[i].value = value
					continue stepLoop
				}
			}
			boxes[h] = append(lenses, lens{label, value})
			continue stepLoop
		}
		for i := 0; i < len(lenses); i++ {
			if lenses[i].label == label {
				copy(lenses[i:], lenses[i+1:])
				boxes[h] = lenses[:len(lenses)-1]
				continue stepLoop
			}
		}
	}
	sum := 0
	for i, lenses := range boxes {
		for j, l := range lenses {
			sum += (i + 1) * (j + 1) * l.value
		}
	}
	fmt.Printf("Day 15, Problem 2, Answer: %v\n", sum)
}
