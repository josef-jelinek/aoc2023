package main

import (
	"slices"
	"strconv"
	"strings"
)

func solveDay22Part1(input string) any {
	bricks := parseBricks(strings.Split(strings.TrimSpace(input), "\n"))
	unders := settleBricks(bricks)
	count := 0
countLoop:
	for i := range bricks {
		for _, us := range unders {
			if len(us) == 1 && us[0] == i {
				continue countLoop
			}
		}
		count++
	}
	return count
}

func solveDay22Part2(input string) any {
	bricks := parseBricks(strings.Split(strings.TrimSpace(input), "\n"))
	unders := settleBricks(bricks)
	// Find which bricks makes other bricks fall when destroyed.
	cascade := make([]bool, len(bricks))
	// Get a list which bricks the given brick supports.
	supports := make([][]int, len(bricks))
	for i := range bricks {
		for j, us := range unders {
			if slices.Contains(us, i) {
				if len(us) == 1 {
					cascade[i] = true
				}
				supports[i] = append(supports[i], j)
			}
		}
	}
	// Count the number bricks which fall when a "cascading" brick is destroyed.
	count := 0
	for i := range bricks {
		if !cascade[i] {
			continue
		}
		destroyed := make(map[int]bool)
		destroyed[i] = true
		var q []int
		for _, k := range supports[i] {
			q = append(q, k)
		}
	processQueue:
		for len(q) > 0 {
		findCollapsing:
			for _, j := range q {
				if destroyed[j] {
					q = q[1:]
					continue processQueue
				}
				for _, k := range unders[j] {
					if !destroyed[k] {
						continue findCollapsing
					}
				}
				count++
				q = q[1:]
				destroyed[j] = true
				for _, k := range supports[j] {
					q = append(q, k)
				}
				continue processQueue
			}
			break // no more bricks in queue to destroy
		}
	}
	return count
}

func parseBricks(lines []string) [][2][3]int {
	var bricks [][2][3]int // 2 "corners" with x, y, z components
	for _, s := range lines {
		u, v, _ := strings.Cut(s, "~")
		us := strings.Split(u, ",")
		vs := strings.Split(v, ",")
		var b [2][3]int
		b[0][0], _ = strconv.Atoi(us[0])
		b[0][1], _ = strconv.Atoi(us[1])
		b[0][2], _ = strconv.Atoi(us[2])
		b[1][0], _ = strconv.Atoi(vs[0])
		b[1][1], _ = strconv.Atoi(vs[1])
		b[1][2], _ = strconv.Atoi(vs[2])
		if b[0][2] > b[1][2] { // u below v
			b[0], b[1] = b[1], b[0]
		}
		bricks = append(bricks, b)
	}
	return bricks
}

// settleBricks return list of bricks supporting a given brick for each brick.
// It modified the given bricks to have compacted z co-ordinates.
func settleBricks(bricks [][2][3]int) [][]int {
	unders := make([][]int, len(bricks))
	settled := make([]bool, len(bricks))
	// Find candidates which are falling under each brick.
	for i, b := range bricks {
		bs := findBricksUnder(b, bricks)
		unders[i] = bs
	}
	// Adjust z for all bricks
	for slices.Contains(settled, false) {
	settle:
		for i := range bricks {
			if settled[i] {
				continue
			}
			for _, c := range unders[i] {
				if !settled[c] {
					continue settle
				}
			}
			// This brick can settle now.
			z := 0
			for _, c := range unders[i] {
				z = max(z, bricks[c][1][2])
			}
			d := bricks[i][1][2] - bricks[i][0][2]
			bricks[i][0][2] = z + 1
			bricks[i][1][2] = z + 1 + d
			settled[i] = true
			break
		}
	}
	// Remove any bricks from the supporting list if they are not right underneath.
	for i, us := range unders {
		var vs []int
		for _, c := range us {
			if bricks[c][1][2]+1 == bricks[i][0][2] {
				vs = append(vs, c)
			}
		}
		unders[i] = vs
	}
	return unders
}

// findBricksUnder picks potential candidates which end up under the given brick
// with overlapping x and y co-ordinates.
func findBricksUnder(b [2][3]int, bs [][2][3]int) []int {
	var cs []int
	for i, c := range bs {
		if b[0][2] <= c[1][2] {
			continue
		}
		if !(min(b[0][0], b[1][0]) <= c[0][0] && c[0][0] <= max(b[0][0], b[1][0]) ||
			min(b[0][0], b[1][0]) <= c[1][0] && c[1][0] <= max(b[0][0], b[1][0]) ||
			min(c[0][0], c[1][0]) <= b[0][0] && b[0][0] <= max(c[0][0], c[1][0]) ||
			min(c[0][0], c[1][0]) <= b[1][0] && b[1][0] <= max(c[0][0], c[1][0])) {
			continue
		}
		if !(min(b[0][1], b[1][1]) <= c[0][1] && c[0][1] <= max(b[0][1], b[1][1]) ||
			min(b[0][1], b[1][1]) <= c[1][1] && c[1][1] <= max(b[0][1], b[1][1]) ||
			min(c[0][1], c[1][1]) <= b[0][1] && b[0][1] <= max(c[0][1], c[1][1]) ||
			min(c[0][1], c[1][1]) <= b[1][1] && b[1][1] <= max(c[0][1], c[1][1])) {
			continue
		}
		cs = append(cs, i)
	}
	return cs
}
