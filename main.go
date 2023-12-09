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
	case "7-1":
		solveDay7Part1(input)
	case "7-2":
		solveDay7Part2(input)
	case "8-1":
		solveDay8Part1(input)
	case "8-2":
		solveDay8Part2(input)
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

func solveDay7Part1(input string) {
	type hand struct {
		cards string
		bid   int
	}

	kinds := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	lines := strings.Split(input, "\n")
	var hands []hand
	for _, s := range lines {
		cards, bidStr, ok := strings.Cut(s, " ")
		if !ok {
			continue
		}
		bid, _ := strconv.Atoi(bidStr)
		hands = append(hands, hand{cards, bid})
	}

	countKinds := func(cards string) []int {
		var counts []int
		for s := cards; s != ""; s = strings.ReplaceAll(s, s[:1], "") {
			counts = append(counts, strings.Count(s, s[:1]))
		}
		slices.Sort(counts)
		slices.Reverse(counts)
		return counts
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if cmp := slices.Compare(countKinds(a.cards), countKinds(b.cards)); cmp != 0 {
			return cmp
		}
		return slices.CompareFunc([]byte(a.cards), []byte(b.cards), func(aa, bb byte) int {
			return slices.Index(kinds, bb) - slices.Index(kinds, aa) // Higher rank -> lower index.
		})
	})
	sum := 0
	for i, v := range hands {
		sum += v.bid * (i + 1)
	}
	fmt.Printf("Day 7, Problem 1, Answer: %v\n", sum)
}

func solveDay7Part2(input string) {
	type hand struct {
		cards string
		bid   int
	}

	kinds := []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	lines := strings.Split(input, "\n")
	var hands []hand
	for _, s := range lines {
		cards, bidStr, ok := strings.Cut(s, " ")
		if !ok {
			continue
		}
		bid, _ := strconv.Atoi(bidStr)
		hands = append(hands, hand{cards, bid})
	}

	countKinds := func(cards string) []int {
		jokers := strings.Count(cards, "J")
		cards = strings.ReplaceAll(cards, "J", "")
		if cards == "" {
			return []int{jokers}
		}
		var counts []int
		for s := cards; s != ""; s = strings.ReplaceAll(s, s[:1], "") {
			counts = append(counts, strings.Count(s, s[:1]))
		}
		slices.Sort(counts)
		slices.Reverse(counts)
		counts[0] += jokers
		return counts
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if cmp := slices.Compare(countKinds(a.cards), countKinds(b.cards)); cmp != 0 {
			return cmp
		}
		return slices.CompareFunc([]byte(a.cards), []byte(b.cards), func(aa, bb byte) int {
			return slices.Index(kinds, bb) - slices.Index(kinds, aa) // Higher rank -> lower index.
		})
	})
	sum := 0
	for i, v := range hands {
		sum += v.bid * (i + 1)
	}
	fmt.Printf("Day 7, Problem 2, Answer: %v\n", sum)
}

func solveDay8Part1(input string) {
	lines := strings.Split(input, "\n")
	dirs := lines[0]
	navMap := make(map[string][]string)
	for _, s := range lines[2:] {
		k, lr, ok := strings.Cut(s, " = ")
		if !ok {
			continue
		}
		navMap[k] = strings.Split(strings.Trim(lr, "()"), ", ")
	}
	pos := "AAA"
	i := 0
	for pos != "ZZZ" {
		if dirs[i%len(dirs)] == 'L' {
			pos = navMap[pos][0]
		} else {
			pos = navMap[pos][1]
		}
		i++
	}
	fmt.Printf("Day 8, Problem 1, Answer: %v\n", i)
}

func solveDay8Part2(input string) {
	lines := strings.Split(input, "\n")
	dirs := lines[0]
	idMap := make(map[string]int)
	var ls, rs []string
	var starts, ends []bool
	var startIDs []int
	for i, s := range lines[2:] {
		k, lr, ok := strings.Cut(s, " = ")
		if !ok {
			continue
		}
		idMap[k] = i
		l, r, _ := strings.Cut(strings.Trim(lr, "()"), ", ")
		ls = append(ls, l)
		rs = append(rs, r)
		isStart := strings.HasSuffix(k, "A")
		if isStart {
			startIDs = append(startIDs, i)
		}
		starts = append(starts, isStart)
		ends = append(ends, strings.HasSuffix(k, "Z"))
	}
	lNext := make([]int, len(ls))
	rNext := make([]int, len(rs))
	for i := range ls {
		lNext[i] = idMap[ls[i]]
		rNext[i] = idMap[rs[i]]
	}
	// Cutting corners with assumptions after analyzing input:
	// - single end location - severe limitation of generality
	// - end location inside the loop - if not true, naive solution would work
	loopStarts := make([]int, len(startIDs))
	loopLens := make([]int, len(startIDs))
	loopEndStates := make([]int, len(startIDs))
	for i, id := range startIDs {
		steps := 0
		visited := make(map[[2]int]int) // steps%len(dirs), steps
		var endStateSteps []int
		for {
			j := steps % len(dirs)
			if k, ok := visited[[2]int{j, id}]; ok {
				loopStarts[i] = k
				loopLens[i] = steps - k
				// Applying assumption of one end state.
				// Take the last to conform to 8-3-sample.txt, which breaks this assumption.
				loopEndStates[i] = endStateSteps[len(endStateSteps)-1] - k
				break
			}
			if ends[id] {
				endStateSteps = append(endStateSteps, steps)
			}
			visited[[2]int{j, id}] = steps
			if dirs[j] == 'L' {
				id = lNext[id]
			} else {
				id = rNext[id]
			}
			steps++
		}
	}
	endState := loopStarts[0] + loopEndStates[0]
	period := loopLens[0]
	for i := 1; i < len(startIDs); i++ {
		es := loopStarts[i] + loopEndStates[i]
		for endState != es {
			for endState < es {
				endState += period
			}
			for es < endState {
				es += loopLens[i]
			}
		}
		gcd := period
		n := loopLens[i]
		for n != 0 {
			gcd, n = n, gcd%n
		}
		period *= loopLens[i] / gcd
	}
	fmt.Printf("Day 8, Problem 2, Answer: %v\n", endState)
}
