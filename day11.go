package main

import (
	"strings"
)

func solveDay11Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var xs, ys []int
	var lockedYs, lockedXs = make([]byte, len(lines)), make([]byte, len(lines[0]))
	for y := 0; y < len(lockedYs); y++ {
		line := lines[y]
		for x := 0; x < len(lockedXs); x++ {
			if line[x] != '.' {
				lockedYs[y] = 1
				lockedXs[x] = 1
				ys = append(ys, y)
				xs = append(xs, x)
			}
		}
	}
	sum := 0
	for i := 0; i < len(xs)-1; i++ {
		xi, yi := xs[i], ys[i]
		for j := i + 1; j < len(xs); j++ {
			xj, yj := xs[j], ys[j]
			dx, dy := max(xi, xj)-min(xi, xj), yj-yi
			if dx < 0 {
				dx = -dx
			}
			for x := min(xi, xj) + 1; x < max(xi, xj); x++ {
				if lockedXs[x] == 0 {
					dx++
				}
			}
			for y := yi + 1; y < yj; y++ {
				if lockedYs[y] == 0 {
					dy++
				}
			}
			sum += dx + dy
		}
	}
	return sum
}

func solveDay11Part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var xs, ys []int
	var lockedYs, lockedXs = make([]byte, len(lines)), make([]byte, len(lines[0]))
	for y := 0; y < len(lockedYs); y++ {
		line := lines[y]
		for x := 0; x < len(lockedXs); x++ {
			if line[x] != '.' {
				lockedYs[y] = 1
				lockedXs[x] = 1
				ys = append(ys, y)
				xs = append(xs, x)
			}
		}
	}
	sum := 0
	for i := 0; i < len(xs)-1; i++ {
		xi, yi := xs[i], ys[i]
		for j := i + 1; j < len(xs); j++ {
			xj, yj := xs[j], ys[j]
			dx, dy := max(xi, xj)-min(xi, xj), yj-yi
			if dx < 0 {
				dx = -dx
			}
			for x := min(xi, xj) + 1; x < max(xi, xj); x++ {
				if lockedXs[x] == 0 {
					dx += 999999
				}
			}
			for y := yi + 1; y < yj; y++ {
				if lockedYs[y] == 0 {
					dy += 999999
				}
			}
			sum += dx + dy
		}
	}
	return sum
}
