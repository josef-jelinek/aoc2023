package main

import (
	"fmt"
	"io"
	"os"
)

var solveFuncByID = map[string]func(string){
	"1-1":  solveDay1Part1,
	"1-2":  solveDay1Part2,
	"2-1":  solveDay2Part1,
	"2-2":  solveDay2Part2,
	"3-1":  solveDay3Part1,
	"3-2":  solveDay3Part2,
	"4-1":  solveDay4Part1,
	"4-2":  solveDay4Part2,
	"5-1":  solveDay5Part1,
	"5-2":  solveDay5Part2,
	"6-1":  solveDay6Part1,
	"6-2":  solveDay6Part2,
	"7-1":  solveDay7Part1,
	"7-2":  solveDay7Part2,
	"8-1":  solveDay8Part1,
	"8-2":  solveDay8Part2,
	"9-1":  solveDay9Part1,
	"9-2":  solveDay9Part2,
	"10-1": solveDay10Part1,
	"10-2": solveDay10Part2,
	"11-1": solveDay11Part1,
	"11-2": solveDay11Part2,
	"12-1": solveDay12Part1,
	"12-2": solveDay12Part2,
	"13-1": solveDay13Part1,
	"13-2": solveDay13Part2,
	"14-1": solveDay14Part1,
	"14-2": solveDay14Part2,
	"15-1": solveDay15Part1,
	"15-2": solveDay15Part2,
	"16-1": solveDay16Part1,
	"16-2": solveDay16Part2,
}

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
	solve := solveFuncByID[id]
	if solve == nil {
		return fmt.Errorf("invalid ID: %q", id)
	}
	solve(string(b))
	return nil
}
