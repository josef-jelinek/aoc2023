package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveDay19Part1(input string) {
	type rule struct {
		cond   func(map[byte]int) bool
		action string
	}

	trueCond := func(map[byte]int) bool {
		return true
	}

	blocks := strings.Split(strings.TrimSpace(input), "\n\n")
	workflows := make(map[string][]rule)
	for _, s := range strings.Split(blocks[0], "\n") {
		name, rulesStr, _ := strings.Cut(strings.TrimSuffix(s, "}"), "{")
		var rules []rule
		for _, ruleStr := range strings.Split(rulesStr, ",") {
			if condStr, action, ok := strings.Cut(ruleStr, ":"); ok {
				id := condStr[0]
				op := condStr[1]
				val, _ := strconv.Atoi(condStr[2:])
				cond := func(m map[byte]int) bool {
					return op == '<' && m[id] < val || op == '>' && m[id] > val
				}
				rules = append(rules, rule{cond, action})
				continue
			}
			rules = append(rules, rule{trueCond, ruleStr})
		}
		workflows[name] = rules
	}

	sum := 0
	for _, s := range strings.Split(blocks[1], "\n") {
		rating := make(map[byte]int)
		for _, ratingStr := range strings.Split(strings.Trim(s, "{}"), ",") {
			val, _ := strconv.Atoi(ratingStr[2:])
			rating[ratingStr[0]] = val
		}
		action := "in"
		for action != "A" && action != "R" {
			rules := workflows[action]
			for _, rule := range rules {
				if rule.cond(rating) {
					action = rule.action
					break
				}
			}
		}
		if action == "A" {
			sum += rating['x'] + rating['m'] + rating['a'] + rating['s']
		}
	}
	fmt.Printf("Day 19, Problem 1, Answer: %v\n", sum)
}

func solveDay19Part2(input string) {
	type rule struct {
		id     byte
		op     byte
		val    int
		action string
	}

	blocks := strings.Split(strings.TrimSpace(input), "\n\n")
	workflows := make(map[string][]rule)
	for _, s := range strings.Split(blocks[0], "\n") {
		name, rulesStr, _ := strings.Cut(strings.TrimSuffix(s, "}"), "{")
		var rules []rule
		for _, ruleStr := range strings.Split(rulesStr, ",") {
			if condStr, action, ok := strings.Cut(ruleStr, ":"); ok {
				val, _ := strconv.Atoi(condStr[2:])
				rules = append(rules, rule{id: condStr[0], op: condStr[1], val: val, action: action})
				continue
			}
			rules = append(rules, rule{action: ruleStr})
		}
		workflows[name] = rules
	}

	// imterval [8]int is 4x (low, high) for 'x', 'm', 'a', 's' IDs.
	var addIntervals func(intervals [][8]int, interval [8]int, action string) [][8]int
	addIntervals = func(intervals [][8]int, interval [8]int, action string) [][8]int {
		if action == "R" {
			return intervals
		}
		if action == "A" {
			return append(intervals, interval)
		}
		for _, rule := range workflows[action] {
			i := strings.IndexByte("xmas", rule.id)
			if i < 0 {
				return addIntervals(intervals, interval, rule.action)
			}
			subInterval := interval
			if rule.op == '<' {
				subInterval[2*i+1] = rule.val - 1
				interval[2*i] = rule.val
			} else {
				subInterval[2*i] = rule.val + 1
				interval[2*i+1] = rule.val
			}
			intervals = addIntervals(intervals, subInterval, rule.action)
		}
		return nil // Unreachable unless the input is invalid.
	}

	intervals := addIntervals(nil, [8]int{1, 4000, 1, 4000, 1, 4000, 1, 4000}, "in")
	sum := 0
	for _, i := range intervals {
		sum += (i[1] - i[0] + 1) * (i[3] - i[2] + 1) * (i[5] - i[4] + 1) * (i[7] - i[6] + 1)
	}
	fmt.Printf("Day 19, Problem 2, Answer: %v\n", sum)
}
