package main

import (
	"slices"
	"strconv"
	"strings"
)

func solveDay18Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var dirs []string
	var counts []int
	for _, s := range lines {
		ss := strings.Split(s, " ")
		dirs = append(dirs, ss[0])
		count, _ := strconv.Atoi(ss[1])
		counts = append(counts, count)
	}
	x, y, minX, minY, maxX, maxY := 0, 0, 0, 0, 0, 0
	for i := range dirs {
		switch dirs[i] {
		case "R":
			x += counts[i]
			maxX = max(maxX, x)
		case "D":
			y += counts[i]
			maxY = max(maxY, y)
		case "L":
			x -= counts[i]
			minX = min(minX, x)
		case "U":
			y -= counts[i]
			minY = min(minY, y)
		}
	}
	w, h := maxX-minX+1, maxY-minY+1
	grid := make([][]byte, h)
	ud := make([][]int8, h)
	for y := range grid {
		grid[y] = make([]byte, w)
		ud[y] = make([]int8, w)
	}
	x, y = -minX, -minY
	for i := range dirs {
		switch dirs[i] {
		case "R":
			for j := 0; j < counts[i]; j++ {
				grid[y][x] = 1
				x++
			}
		case "D":
			ud[y][x] = 1
			for j := 0; j < counts[i]; j++ {
				grid[y][x] = 1
				y++
				ud[y][x] = 1
			}
		case "L":
			for j := 0; j < counts[i]; j++ {
				grid[y][x] = 1
				x--
			}
		case "U":
			ud[y][x] = -1
			for j := 0; j < counts[i]; j++ {
				grid[y][x] = 1
				y--
				ud[y][x] = -1
			}
		}
	}
	for y := 0; y < h; y++ {
		x := 0
		for x < w && ud[y][x] == 0 {
			x++
		}
		if x < w {
			d := ud[y][x]
			for x < w {
				for x < w && ud[y][x] != -d {
					grid[y][x] = 1
					x++
				}
				for x < w && ud[y][x] != d {
					x++
				}
			}
		}
	}
	volume := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] > 0 {
				volume++
			}
		}
	}
	return volume
}

func solveDay18Part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var dirs []byte
	var counts []int
	for _, s := range lines {
		ss := strings.Split(s, " ")
		dirs = append(dirs, "RDLU"[ss[2][7]-'0'])
		count, _ := strconv.ParseInt(ss[2][2:7], 16, 64)
		counts = append(counts, int(count))
	}
	x, y := 0, 0
	xs := []int{x}
	ys := []int{y}
	for i := range dirs {
		switch dirs[i] {
		case 'R':
			x += counts[i]
			xs = append(xs, x)
		case 'D':
			y += counts[i]
			ys = append(ys, y)
		case 'L':
			x -= counts[i]
			xs = append(xs, x)
		case 'U':
			y -= counts[i]
			ys = append(ys, y)
		}
	}
	// Set up a grid of unique x and y co-ordinates.
	slices.Sort(xs)
	slices.Sort(ys)
	xs = slices.Compact(xs)
	ys = slices.Compact(ys)
	xMap := make(map[int]int, len(xs))
	yMap := make(map[int]int, len(ys))
	for i, v := range xs {
		xMap[v] = i
	}
	for i, v := range ys {
		yMap[v] = i
	}
	w, h := len(xs), len(ys)
	ud := make([][]int8, h)
	for y := range ud {
		ud[y] = make([]int8, w)
	}
	// Measure circumference, since it is 1-width outline contributing to the volume.
	length := 0
	// Go through the outline and mark up and down movements.
	x, y = 0, 0
	for i := range dirs {
		switch dirs[i] {
		case 'R':
			x += counts[i]
		case 'D':
			n := yMap[y+counts[i]] - yMap[y]
			for j := 0; j < n; j++ {
				ud[yMap[y]+j][xMap[x]] = 1
			}
			y += counts[i]
		case 'L':
			x -= counts[i]
		case 'U':
			n := yMap[y] - yMap[y-counts[i]]
			for j := 0; j < n; j++ {
				ud[yMap[y]-j-1][xMap[x]] = -1
			}
			y -= counts[i]
		}
		length += counts[i]
	}
	volume := 0
	// Go through up/down edges dividing inside and outside and accumulate inside boxes.
	for y := 0; y+1 < h; y++ {
		x := 0
		for x+1 < w {
			for x+1 < w && ud[y][x] == 0 {
				x++
			}
			d := ud[y][x]
			for x+1 < w && ud[y][x] != -d {
				v := (xs[x+1] - xs[x]) * (ys[y+1] - ys[y])
				volume += v
				x++
			}
			x++
		}
	}
	// Increase volume by outside half of the outline and 4 "corners" (convex turns - concave turns).
	volume += length/2 + 1
	return volume
}
