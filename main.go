package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Expected problem ID (e.g \"1-1\", \"1-2\", \"2-1\", ...) and file name\n")
		os.Exit(1)
	}
	var err error
	switch os.Args[1] {
	case "1-1":
		err = solveDay1Part1(os.Args[2])
	case "1-2":
		err = solveDay1Part2(os.Args[2])
	case "2-1":
		err = solveDay2Part1(os.Args[2])
	case "2-2":
		err = solveDay2Part2(os.Args[2])
	case "3-1":
		err = solveDay3Part1(os.Args[2])
	case "3-2":
		err = solveDay3Part2(os.Args[2])
	case "4-1":
		err = solveDay4Part1(os.Args[2])
	case "4-2":
		err = solveDay4Part2(os.Args[2])
	default:
		fmt.Printf("Problem ID not valid: %q\n", os.Args[1])
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func solveDay1Part1(filename string) error {
	fmt.Println("Day 1, Problem 1")
	data, err := fileAsString(filename)
	if err != nil {
		return err
	}
	sum := 0
	for _, s := range strings.Split(data, "\n") {
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
	fmt.Printf("Answer: %d\n", sum)
	return nil
}

func solveDay1Part2(filename string) error {
	fmt.Println("Day 1, Problem 2")
	data, err := fileAsString(filename)
	if err != nil {
		return err
	}
	sum := 0
	for _, s := range strings.Split(data, "\n") {
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
	fmt.Printf("Answer: %d\n", sum)
	return nil
}

func solveDay2Part1(filename string) error {
	fmt.Println("Day 2, Problem 1")
	data, err := fileAsString(filename)
	if err != nil {
		return err
	}
	maxByColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0
gameLoop:
	for _, s := range strings.Split(data, "\n") {
		name, spec, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		id, err := strconv.Atoi(name[5:])
		if err != nil {
			return err
		}
		for _, draw := range strings.Split(spec, "; ") {
			for _, cube := range strings.Split(draw, ", ") {
				count, color, _ := strings.Cut(cube, " ")
				n, err := strconv.Atoi(count)
				if err != nil {
					return err
				}
				if n > maxByColor[color] {
					continue gameLoop
				}
			}
		}
		sum += id
	}
	fmt.Printf("Answer: %d\n", sum)
	return nil
}

func solveDay2Part2(filename string) error {
	fmt.Println("Day 2, Problem 2")
	data, err := fileAsString(filename)
	if err != nil {
		return err
	}
	sum := 0
	for _, s := range strings.Split(data, "\n") {
		_, spec, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		minByColor := make(map[string]int)
		for _, draw := range strings.Split(spec, "; ") {
			for _, cube := range strings.Split(draw, ", ") {
				count, color, _ := strings.Cut(cube, " ")
				n, err := strconv.Atoi(count)
				if err != nil {
					return err
				}
				minByColor[color] = max(minByColor[color], n)
			}
		}
		power := 1
		for _, v := range minByColor {
			power *= v
		}
		sum += power
	}
	fmt.Printf("Answer: %d\n", sum)
	return nil
}

func solveDay3Part1(filename string) error {
	fmt.Println("Day 3, Problem 1")
	data, err := fileAsString(filename)
	if err != nil {
		return err
	}
	lines := strings.Split(data, "\n")

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
	fmt.Printf("Answer: %d\n", sum)
	return nil
}

func solveDay3Part2(filename string) error {
	fmt.Println("Day 3, Problem 2")
	data, err := fileAsString(filename)
	if err != nil {
		return err
	}
	lines := strings.Split(data, "\n")

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
	fmt.Printf("Answer: %d\n", sum)
	return nil
}

func solveDay4Part1(filename string) error {
	fmt.Println("Day 4, Problem 1")
	data, err := fileAsString(filename)
	if err != nil {
		return err
	}
	for _, s := range strings.Split(data, "\n") {
		fmt.Println(s)
	}
	fmt.Printf("Answer: %d\n", 0)
	return nil
}

func solveDay4Part2(filename string) error {
	fmt.Println("Day 4, Problem 2")
	data, err := fileAsString(filename)
	if err != nil {
		return err
	}
	for _, s := range strings.Split(data, "\n") {
		fmt.Println(s)
	}
	fmt.Printf("Answer: %d\n", 0)
	return nil
}

func fileAsString(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
