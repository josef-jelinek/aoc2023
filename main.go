package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Expected problem ID (e.g \"1-1\", \"1-2\", \"2-1\", ...) and file name\n")
		os.Exit(1)
	}
	if err := solve(os.Args[1], os.Args[2]); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func solve(id, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	input := string(b)
	switch os.Args[1] {
	case "1-1":
		solveDay1Part1(input)
	case "1-2":
		solveDay1Part2(input)
	case "2-1":
		solveDay2Part1(input)
	case "2-2":
		solveDay2Part2(input)
	case "3-1":
		solveDay3Part1(input)
	case "3-2":
		solveDay3Part2(input)
	case "4-1":
		solveDay4Part1(input)
	case "4-2":
		solveDay4Part2(input)
	case "5-1":
		solveDay5Part1(input)
	case "5-2":
		solveDay5Part2(input)
	case "6-1":
		solveDay6Part1(input)
	case "6-2":
		solveDay6Part2(input)
	case "7-1":
		solveDay7Part1(input)
	case "7-2":
		solveDay7Part2(input)
	case "8-1":
		solveDay8Part1(input)
	case "8-2":
		solveDay8Part2(input)
	case "9-1":
		solveDay9Part1(input)
	case "9-2":
		solveDay9Part2(input)
	default:
		return fmt.Errorf("invalid ID: %q", id)
	}
	return nil
}
