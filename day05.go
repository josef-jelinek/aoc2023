package main

import (
	"slices"
	"strconv"
	"strings"
)

func solveDay5Part1(input string) any {
	lines := strings.Split(input, "\n")
	seedsStr, _ := strings.CutPrefix(lines[0], "seeds: ")
	var seeds []int
	for _, s := range strings.Split(seedsStr, " ") {
		v, _ := strconv.Atoi(s)
		seeds = append(seeds, v)
	}
	var srcs, dsts []string
	var rngs [][][3]int // dst start, src start, count
	i := 2
	for i < len(lines) {
		mapName, _ := strings.CutSuffix(lines[i], " map:")
		src, dst, _ := strings.Cut(mapName, "-to-")
		srcs = append(srcs, src)
		dsts = append(dsts, dst)
		rngs = append(rngs, nil)
		i++
		for i < len(lines) && lines[i] != "" {
			nums := strings.Split(lines[i], " ")
			var rng [3]int
			rng[0], _ = strconv.Atoi(nums[0])
			rng[1], _ = strconv.Atoi(nums[1])
			rng[2], _ = strconv.Atoi(nums[2])
			rngs[len(rngs)-1] = append(rngs[len(rngs)-1], rng)
			i++
		}
		i++
	}
	var minLocation int
	for i, seed := range seeds {
		src := "seed"
		id := seed
		for src != "location" {
			index := slices.Index(srcs, src)
			for _, rng := range rngs[index] {
				if id >= rng[1] && id < rng[1]+rng[2] {
					id = rng[0] + (id - rng[1])
					break
				}
			}
			src = dsts[index]
		}
		if i == 0 {
			minLocation = id
		}
		minLocation = min(minLocation, id)
	}
	return minLocation
}

func solveDay5Part2(input string) any {
	lines := strings.Split(input, "\n")
	seedsStr, _ := strings.CutPrefix(lines[0], "seeds: ")
	var seeds []int // start1 count1 start2 count2 ...
	for _, s := range strings.Split(seedsStr, " ") {
		v, _ := strconv.Atoi(s)
		seeds = append(seeds, v)
	}
	var srcs, dsts []string
	var rngs [][][3]int // dst start, src start, count
	i := 2
	for i < len(lines) {
		mapName, _ := strings.CutSuffix(lines[i], " map:")
		src, dst, _ := strings.Cut(mapName, "-to-")
		srcs = append(srcs, src)
		dsts = append(dsts, dst)
		rngs = append(rngs, nil)
		i++
		for i < len(lines) && lines[i] != "" {
			nums := strings.Split(lines[i], " ")
			var rng [3]int
			rng[0], _ = strconv.Atoi(nums[0])
			rng[1], _ = strconv.Atoi(nums[1])
			rng[2], _ = strconv.Atoi(nums[2])
			rngs[len(rngs)-1] = append(rngs[len(rngs)-1], rng)
			i++
		}
		i++
	}
	minLocation := -1
	for i := 0; i < len(seeds); i += 2 {
		seedID := seeds[i]
		for seedID < seeds[i]+seeds[i+1] {
			src := "seed"
			id := seedID
			// This is the first trickier optimization.
			// Since we are working with ranges, minimum ID will always be the
			// lower bound and higher IDs do not need checking as long as we
			// stay inside the interval intersection.
			skip := seeds[i] + seeds[i+1] - seedID
			for src != "location" {
				index := slices.Index(srcs, src)
				for _, rng := range rngs[index] {
					if id >= rng[1] && id < rng[1]+rng[2] {
						id = rng[0] + (id - rng[1])
						skip = min(skip, rng[0]+rng[2]-id)
						break
					}
				}
				src = dsts[index]
			}
			if minLocation < 0 {
				minLocation = id
			} else {
				minLocation = min(minLocation, id)
			}
			seedID += skip
		}
	}
	return minLocation
}
