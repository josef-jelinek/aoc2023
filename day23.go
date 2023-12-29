package main

import (
	"strings"
)

func solveDay23Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return maxPath(lines, 1, 0, 0, make(map[[2]int]bool))
}

func solveDay23Part2(input string) any {
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	w, h := len(lines[0]), len(lines)
	// Find "crossroads" as unique graph nodes.
	var nodes [][2]int
	nodes = append(nodes, [2]int{1, 0}, [2]int{w - 2, h - 1})
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '#' {
				continue
			}
			exits := 0
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < w && ny >= 0 && ny < h && lines[ny][nx] != '#' {
					exits++
				}
			}
			if exits >= 3 {
				nodes = append(nodes, [2]int{x, y})
			}
		}
	}
	// Graph is stored as both martix of nodes for a node and node-node distances.
	nodeIDs := make(map[[2]int]int)
	edges := make([][]int, len(nodes))
	dists := make([][]int, len(nodes))
	for i, xy := range nodes {
		nodeIDs[xy] = i
		dists[i] = make([]int, len(nodes))
	}
	for id := range nodes {
		var q [][3]int
		visited := make(map[[2]int]bool)
		x, y := nodes[id][0], nodes[id][1]
		q = append(q, [3]int{x, y, 0})
		visited[[2]int{x, y}] = true
		for len(q) > 0 {
			x, y, d := q[0][0], q[0][1], q[0][2]
			q = q[1:]
			if d > 0 {
				if i, found := nodeIDs[[2]int{x, y}]; found && i != id {
					edges[id] = append(edges[id], i)
					dists[id][i] = d
					continue
				}
			}
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < w && ny >= 0 && ny < h && lines[ny][nx] != '#' && !visited[[2]int{nx, ny}] {
					q = append(q, [3]int{nx, ny, d + 1})
					visited[[2]int{nx, ny}] = true
				}
			}
		}
	}
	return maxGraphPath(edges, dists, 0, 0, make(map[int]bool))
}

func maxPath(lines []string, x, y, dist int, visited map[[2]int]bool) int {
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	w, h := len(lines[0]), len(lines)
	if x == w-2 && y == h-1 {
		return dist
	}
	visited[[2]int{x, y}] = true
	diri := strings.IndexByte(">v<^", lines[y][x])
	maxd := 0
	for i, dir := range dirs {
		if diri >= 0 && i != diri {
			continue
		}
		nx, ny := x+dir[0], y+dir[1]
		if nx < 0 || nx >= w || ny < 0 || ny >= h || lines[ny][nx] == '#' || visited[[2]int{nx, ny}] {
			continue
		}
		maxd = max(maxd, maxPath(lines, nx, ny, dist+1, visited))
	}
	visited[[2]int{x, y}] = false
	return maxd
}

func maxGraphPath(edges [][]int, dists [][]int, id, dist int, visited map[int]bool) int {
	if id == 1 {
		return dist
	}
	visited[id] = true
	maxd := 0
	for _, i := range edges[id] {
		if visited[i] {
			continue
		}
		maxd = max(maxd, maxGraphPath(edges, dists, i, dist+dists[id][i], visited))
	}
	visited[id] = false
	return maxd
}
