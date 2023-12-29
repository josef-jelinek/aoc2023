package main

import (
	"strings"
)

func solveDay25Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	graph := make(map[string]map[string]bool)
	nodes := make(map[string]bool)
	for _, s := range lines {
		u, vs, _ := strings.Cut(s, ": ")
		if graph[u] == nil {
			graph[u] = make(map[string]bool)
			nodes[u] = true
		}
		for _, v := range strings.Split(vs, " ") {
			graph[u][v] = true
			if graph[v] == nil {
				graph[v] = make(map[string]bool)
				nodes[v] = true
			}
			graph[v][u] = true
		}
	}
	for {
		maxNode := ""
		maxXconns := 0
		sumXconns := 0
		for n := range nodes {
			xconns := 0
			for m := range graph[n] {
				if !nodes[m] {
					xconns++
					sumXconns++
				}
			}
			if maxXconns == 0 || maxXconns < xconns {
				maxNode = n
				maxXconns = xconns
			}
		}
		if sumXconns == 3 {
			break
		}
		delete(nodes, maxNode)
	}
	return len(nodes) * (len(graph) - len(nodes))
}
