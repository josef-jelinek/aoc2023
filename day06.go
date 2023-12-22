package main

import (
	"strconv"
	"strings"
)

func solveDay6Part1(input string) any {
	lines := strings.Split(input, "\n")
	var times, dists []int
	timesStr, _ := strings.CutPrefix(lines[0], "Time:")
	distsStr, _ := strings.CutPrefix(lines[1], "Distance:")
	for _, v := range strings.Fields(timesStr) {
		t, _ := strconv.Atoi(v)
		times = append(times, t)
	}
	for _, v := range strings.Fields(distsStr) {
		d, _ := strconv.Atoi(v)
		dists = append(dists, d)
	}
	product := 1
	for i, t := range times {
		// Symetrical expr. so finding start means finding end as well.
		for j := 1; j < (t+1)/2; j++ {
			d := j * (t - j)
			if d > dists[i] {
				product *= t - 2*j + 1
				break
			}
		}
	}
	return product
}

func solveDay6Part2(input string) any {
	lines := strings.Split(input, "\n")
	timeStr, _ := strings.CutPrefix(lines[0], "Time:")
	distStr, _ := strings.CutPrefix(lines[1], "Distance:")
	time, _ := strconv.Atoi(strings.Join(strings.Fields(timeStr), ""))
	dist, _ := strconv.Atoi(strings.Join(strings.Fields(distStr), ""))
	var count int
	// If slow, binary search can be used as the search is on a half parabola.
	// (Or even something faster as a Newton's method.)
	for i := 1; i < (time+1)/2; i++ {
		d := i * (time - i)
		if d > dist {
			count = time - 2*i + 1
			break
		}
	}
	return count
}
