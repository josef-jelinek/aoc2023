package main

import (
	"slices"
	"strconv"
	"strings"
)

func solveDay7Part1(input string) any {
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
	return sum
}

func solveDay7Part2(input string) any {
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
	return sum
}
