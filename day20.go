package main

import (
	"slices"
	"strings"
)

func solveDay20Part1(input string) any {
	type signal struct {
		src, dst int
		lev      byte
	}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	n := len(lines)
	var idByName = make(map[string]int)
	var names []string
	var kinds []byte
	var dstNames [][]string
	for i, s := range lines {
		src, dsts, _ := strings.Cut(s, " -> ")
		name := strings.TrimLeft(src, "%&")
		idByName[name] = i
		names = append(names, name)
		kinds = append(kinds, src[0])
		dstNames = append(dstNames, strings.Split(dsts, ", "))
	}
	srcIDs := make([][]int, n)
	dstIDs := make([][]int, n)
	for i, dsts := range dstNames {
		for _, dst := range dsts {
			id, found := idByName[dst]
			if !found {
				id = len(names)
				names = append(names, dst)
			} else {
				srcIDs[id] = append(srcIDs[id], i)
			}
			dstIDs[i] = append(dstIDs[i], id)
		}
	}
	states := make([][]byte, n)
	for i := range states {
		switch kinds[i] {
		case '%':
			states[i] = []byte{0}
		case '&':
			states[i] = make([]byte, len(srcIDs[i]))
		}
	}
	id0 := idByName["broadcaster"]
	var loopLevs [][2]int
loop:
	for {
		qSignal := []signal{{src: -1, dst: id0, lev: 0}} // button -low-> broadcaster
		levs := [2]int{1, 0}                             // button low signal
		for len(qSignal) > 0 {
			s := qSignal[0]
			qSignal = qSignal[1:]
			id := s.dst
			if id >= len(kinds) {
				continue
			}
			switch kinds[id] {
			case 'b':
				for _, dst := range dstIDs[id] {
					qSignal = append(qSignal, signal{src: id, dst: dst, lev: s.lev})
					levs[s.lev]++
				}
			case '%':
				if s.lev == 0 {
					states[id][0] = 1 - states[id][0]
					for _, dst := range dstIDs[id] {
						qSignal = append(qSignal, signal{src: id, dst: dst, lev: states[id][0]})
						levs[states[id][0]]++
					}
				}
			case '&':
				var lev byte = 1
				for i, src := range srcIDs[id] {
					if s.src == src {
						states[id][i] = s.lev
					}
					lev &= states[id][i]
				}
				for _, dst := range dstIDs[id] {
					qSignal = append(qSignal, signal{src: id, dst: dst, lev: 1 - lev})
					levs[1-lev]++
				}
			}
		}
		loopLevs = append(loopLevs, levs)
		if len(loopLevs) == 1000 {
			break
		}
		// Test whether we are back to the default state.
		for _, state := range states {
			for _, v := range state {
				if v != 0 {
					continue loop
				}
			}
		}
		break
	}
	// Seems like there is no return to the reset state, so this is redundant.
	// It is left here though as the small examples need it.
	d := 1000 / len(loopLevs)
	m := 1000 % len(loopLevs)
	var lo, hi int
	for i, levs := range loopLevs {
		lo += d * levs[0]
		hi += d * levs[1]
		if i < m {
			lo += levs[0]
			hi += levs[1]
		}
	}
	return lo * hi
}

func solveDay20Part2(input string) any {
	type signal struct {
		src, dst int
		lev      byte
	}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var idByName = make(map[string]int)
	var names []string
	var kinds []byte
	var dstNames [][]string
	for i, s := range lines {
		src, dsts, _ := strings.Cut(s, " -> ")
		name := strings.TrimLeft(src, "%&")
		idByName[name] = i
		names = append(names, name)
		kinds = append(kinds, src[0])
		dstNames = append(dstNames, strings.Split(dsts, ", "))
	}
	n := len(lines)
	// Add any names, which are not in source names.
	for _, dsts := range dstNames {
		for _, dst := range dsts {
			id, found := idByName[dst]
			if !found {
				id = n
				idByName[dst] = id
				names = append(names, dst)
				kinds = append(kinds, '?')
				n++
			}
		}
	}
	srcIDs := make([][]int, n)
	dstIDs := make([][]int, n)
	for i, dsts := range dstNames {
		for _, dst := range dsts {
			id := idByName[dst]
			srcIDs[id] = append(srcIDs[id], i)
			dstIDs[i] = append(dstIDs[i], id)
		}
	}

	states := make([][]byte, n)
	for i := range states {
		switch kinds[i] {
		case '%':
			states[i] = []byte{0}
		case '&':
			states[i] = make([]byte, len(srcIDs[i]))
		}
	}

	// Find out which conjunction to observe for all high inputs.
	idEnd := idByName["rx"]
	if len(srcIDs[idEnd]) != 1 || kinds[srcIDs[idEnd][0]] != '&' {
		panic("Unsupported configuration")
	}
	idEnd = srcIDs[idEnd][0] // watch inputs to this conjunction
	id0 := idByName["broadcaster"]
	firstHighs := make([]int, len(srcIDs[idEnd]))
	loopCount := 0
loop:
	for {
		loopCount++
		qSignal := []signal{{src: -1, dst: id0, lev: 0}} // button -low-> broadcaster
		for len(qSignal) > 0 {
			s := qSignal[0]
			qSignal = qSignal[1:]
			id := s.dst
			// Test whether we see a high pulse to an end conjunction input.
			if id == idEnd && s.lev == 1 {
				i := slices.Index(srcIDs[idEnd], s.src)
				if firstHighs[i] == 0 {
					firstHighs[i] = loopCount
				}
				// Find out whether we got first high pulses for all inputs.
				if slices.Min(firstHighs) > 0 {
					break loop
				}
			}
			switch kinds[id] {
			case 'b':
				for _, dst := range dstIDs[id] {
					qSignal = append(qSignal, signal{src: id, dst: dst, lev: s.lev})
				}
			case '%':
				if s.lev == 0 {
					states[id][0] = 1 - states[id][0]
					for _, dst := range dstIDs[id] {
						qSignal = append(qSignal, signal{src: id, dst: dst, lev: states[id][0]})
					}
				}
			case '&':
				var lev byte = 1
				for i, src := range srcIDs[id] {
					if s.src == src {
						states[id][i] = s.lev
					}
					lev &= states[id][i]
				}
				for _, dst := range dstIDs[id] {
					qSignal = append(qSignal, signal{src: id, dst: dst, lev: 1 - lev})
				}
			}
		}
	}
	// Seems like least common multiple is redundant.
	firstHigh := firstHighs[0]
	for i := 1; i < len(firstHighs); i++ {
		gcd, n := firstHigh, firstHighs[i]
		for n != 0 {
			gcd, n = n, gcd%n
		}
		firstHigh *= firstHighs[i] / gcd
	}
	return firstHigh
}
