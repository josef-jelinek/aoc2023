package main

import (
	"strings"
)

func solveDay3Part1(input string) any {
	lines := strings.Split(input, "\n")

	isDigit := func(row, col int) bool {
		if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) {
			return false
		}
		b := lines[row][col]
		return '0' <= b && b <= '9'
	}

	isSym := func(row, col int) bool {
		if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) {
			return false
		}
		b := lines[row][col]
		return (b < '0' || '9' < b) && b != '.'
	}

	sum := 0
	for row := range lines {
		inNum := false
		adjSym := false
		num := 0
		for col := 0; col <= len(lines[row]); col++ {
			if isDigit(row, col) {
				if !inNum && (isSym(row-1, col-1) || isSym(row, col-1) || isSym(row+1, col-1)) {
					adjSym = true
				}
				if !adjSym && (isSym(row-1, col) || isSym(row+1, col)) {
					adjSym = true
				}
				num = 10*num + int(lines[row][col]-'0')
				inNum = true
			} else if inNum {
				if adjSym || isSym(row-1, col) || isSym(row, col) || isSym(row+1, col) {
					sum += num
					adjSym = false
				}
				num = 0
				inNum = false
			}
		}
	}
	return sum
}

func solveDay3Part2(input string) any {
	lines := strings.Split(input, "\n")

	isDigit := func(row, col int) bool {
		if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) {
			return false
		}
		b := lines[row][col]
		return '0' <= b && b <= '9'
	}

	isGear := func(row, col int) bool {
		if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) {
			return false
		}
		return lines[row][col] == '*'
	}

	gearNums := make(map[[2]int][]int)
	for row := range lines {
		inNum := false
		num := 0
		var gears [][2]int
		for col := 0; col <= len(lines[row]); col++ {
			if isDigit(row, col) {
				if !inNum {
					if isGear(row-1, col-1) {
						gears = append(gears, [2]int{row - 1, col - 1})
					}
					if isGear(row, col-1) {
						gears = append(gears, [2]int{row, col - 1})
					}
					if isGear(row+1, col-1) {
						gears = append(gears, [2]int{row + 1, col - 1})
					}
				}
				if isGear(row-1, col) {
					gears = append(gears, [2]int{row - 1, col})
				}
				if isGear(row+1, col) {
					gears = append(gears, [2]int{row + 1, col})
				}
				num = 10*num + int(lines[row][col]-'0')
				inNum = true
			} else if inNum {
				if isGear(row-1, col) {
					gears = append(gears, [2]int{row - 1, col})
				}
				if isGear(row, col) {
					gears = append(gears, [2]int{row, col})
				}
				if isGear(row+1, col) {
					gears = append(gears, [2]int{row + 1, col})
				}
				for _, gear := range gears {
					gearNums[gear] = append(gearNums[gear], num)
				}
				gears = gears[:0]
				num = 0
				inNum = false
			}
		}
	}
	sum := 0
	for _, vs := range gearNums {
		if len(vs) != 2 {
			continue
		}
		sum += vs[0] * vs[1]
	}
	return sum
}
