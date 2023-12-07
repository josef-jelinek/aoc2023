package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
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
	default:
		return fmt.Errorf("invalid ID: %q", id)
	}
	return nil
}

func solveDay1Part1(input string) {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		var d10, d1 int
		d10Set := false
		for _, r := range s {
			if '0' <= r && r <= '9' {
				d1 = int(r - '0')
				if !d10Set {
					d10 = 10 * d1
					d10Set = true
				}
			}
		}
		sum += d10 + d1
	}
	fmt.Printf("Day 1, Problem 1, Answer: %d\n", sum)
}

func solveDay1Part2(input string) {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		var d10, d1 int
		d10Set := false
		for i, r := range s {
			d := 0
			switch {
			case '0' <= r && r <= '9':
				d = int(r - '0')
			case strings.HasPrefix(s[i:], "one"):
				d = 1
			case strings.HasPrefix(s[i:], "two"):
				d = 2
			case strings.HasPrefix(s[i:], "three"):
				d = 3
			case strings.HasPrefix(s[i:], "four"):
				d = 4
			case strings.HasPrefix(s[i:], "five"):
				d = 5
			case strings.HasPrefix(s[i:], "six"):
				d = 6
			case strings.HasPrefix(s[i:], "seven"):
				d = 7
			case strings.HasPrefix(s[i:], "eight"):
				d = 8
			case strings.HasPrefix(s[i:], "nine"):
				d = 9
			default:
				continue
			}
			d1 = d
			if !d10Set {
				d10 = 10 * d
				d10Set = true
			}
		}
		sum += d10 + d1
	}
	fmt.Printf("Day 1, Problem 2, Answer: %d\n", sum)
}

func solveDay2Part1(input string) {
	maxByColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0
gameLoop:
	for _, s := range strings.Split(input, "\n") {
		name, spec, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		id, _ := strconv.Atoi(name[5:])
		for _, draw := range strings.Split(spec, "; ") {
			for _, cube := range strings.Split(draw, ", ") {
				count, color, _ := strings.Cut(cube, " ")
				n, _ := strconv.Atoi(count)
				if n > maxByColor[color] {
					continue gameLoop
				}
			}
		}
		sum += id
	}
	fmt.Println("Day 2, Problem 1, Answer: %d\n", sum)
}

func solveDay2Part2(input string) {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		_, spec, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		minByColor := make(map[string]int)
		for _, draw := range strings.Split(spec, "; ") {
			for _, cube := range strings.Split(draw, ", ") {
				count, color, _ := strings.Cut(cube, " ")
				n, _ := strconv.Atoi(count)
				minByColor[color] = max(minByColor[color], n)
			}
		}
		power := 1
		for _, v := range minByColor {
			power *= v
		}
		sum += power
	}
	fmt.Printf("Day 2, Problem 2, Answer: %d\n", sum)
}

func solveDay3Part1(input string) {
	lines := strings.Split(input, "\n")

	isDigit := func(row, col int) bool {
		if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) {
			return false
		}
		b := lines[row][col]
		return '0' <= b && b <= '9'
	}

	isSym := func(row, col int) bool {
		if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) {
			return false
		}
		b := lines[row][col]
		return (b < '0' || '9' < b) && b != '.'
	}

	sum := 0
	for row := range lines {
		inNum := false
		adjSym := false
		num := 0
		for col := 0; col <= len(lines[row]); col++ {
			if isDigit(row, col) {
				if !inNum && (isSym(row-1, col-1) || isSym(row, col-1) || isSym(row+1, col-1)) {
					adjSym = true
				}
				if !adjSym && (isSym(row-1, col) || isSym(row+1, col)) {
					adjSym = true
				}
				num = 10*num + int(lines[row][col]-'0')
				inNum = true
			} else if inNum {
				if adjSym || isSym(row-1, col) || isSym(row, col) || isSym(row+1, col) {
					sum += num
					adjSym = false
				}
				num = 0
				inNum = false
			}
		}
	}
	fmt.Printf("Day 3, Problem 1, Answer: %d\n", sum)
}

func solveDay3Part2(input string) {
	lines := strings.Split(input, "\n")

	isDigit := func(row, col int) bool {
		if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) {
			return false
		}
		b := lines[row][col]
		return '0' <= b && b <= '9'
	}

	isGear := func(row, col int) bool {
		if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[row]) {
			return false
		}
		return lines[row][col] == '*'
	}

	gearNums := make(map[[2]int][]int)
	for row := range lines {
		inNum := false
		num := 0
		var gears [][2]int
		for col := 0; col <= len(lines[row]); col++ {
			if isDigit(row, col) {
				if !inNum {
					if isGear(row-1, col-1) {
						gears = append(gears, [2]int{row - 1, col - 1})
					}
					if isGear(row, col-1) {
						gears = append(gears, [2]int{row, col - 1})
					}
					if isGear(row+1, col-1) {
						gears = append(gears, [2]int{row + 1, col - 1})
					}
				}
				if isGear(row-1, col) {
					gears = append(gears, [2]int{row - 1, col})
				}
				if isGear(row+1, col) {
					gears = append(gears, [2]int{row + 1, col})
				}
				num = 10*num + int(lines[row][col]-'0')
				inNum = true
			} else if inNum {
				if isGear(row-1, col) {
					gears = append(gears, [2]int{row - 1, col})
				}
				if isGear(row, col) {
					gears = append(gears, [2]int{row, col})
				}
				if isGear(row+1, col) {
					gears = append(gears, [2]int{row + 1, col})
				}
				for _, gear := range gears {
					gearNums[gear] = append(gearNums[gear], num)
				}
				gears = gears[:0]
				num = 0
				inNum = false
			}
		}
	}
	sum := 0
	for _, vs := range gearNums {
		if len(vs) != 2 {
			continue
		}
		sum += vs[0] * vs[1]
	}
	fmt.Printf("Day 3, Problem 2, Answer: %d\n", sum)
}

func solveDay4Part1(input string) {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		_, numsPart, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		wantPart, gotPart, _ := strings.Cut(numsPart, " | ")
		wantNums := strings.Split(wantPart, " ")
		gotNums := strings.Split(gotPart, " ")
		count := 0
		for _, gotNum := range gotNums {
			if gotNum != "" && slices.Contains(wantNums, gotNum) {
				count *= 2
				if count == 0 {
					count = 1
				}
			}
		}
		sum += count
	}
	fmt.Println("Day 4, Problem 1, Answer: %d\n", sum)
}

func solveDay4Part2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	won := make([]int, len(lines))
	for i, s := range lines {
		_, numsStr, ok := strings.Cut(s, ": ")
		if !ok {
			continue
		}
		wantNumsStr, gotNumsStr, _ := strings.Cut(numsStr, " | ")
		wantNums := strings.Split(wantNumsStr, " ")
		gotNums := strings.Split(gotNumsStr, " ")
		count := 0
		for _, gotNum := range gotNums {
			if gotNum != "" && slices.Contains(wantNums, gotNum) {
				count++
			}
		}
		for j := 0; j < count; j++ {
			won[i+j+1] += 1 + won[i]
		}
	}
	sum := 0
	for _, w := range won {
		sum += 1 + w
	}
	fmt.Printf("Day 4, Problem 2, Answer: %d\n", sum)
}

func solveDay5Part1(input string) {
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
	fmt.Printf("Day 5, Problem 1, Answer: %d\n", minLocation)
}

func solveDay5Part2(input string) {
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
	fmt.Printf("Day 5, Problem 2, Answer: %d\n", minLocation)
}

func solveDay6Part1(input string) {
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
	fmt.Printf("Day 6, Problem 1, Answer: %d\n", product)
}

func solveDay6Part2(input string) {
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
	fmt.Printf("Day 6, Problem 2, Answer: %d\n", count)
}
