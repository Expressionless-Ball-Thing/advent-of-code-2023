package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	hand string
	bid int
	handtype int
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func determineHandType(stuff string, bid string, joker bool) hand {

	bidnum, _ := strconv.ParseInt(regexp.MustCompile(`\d+`).FindString(bid), 0, 64)
	hand := hand {
		stuff, int(bidnum), 0,
	}

	var freq = make(map[rune]int)

	for _, i := range stuff {
		_, contains := freq[i]
		if !contains {
			freq[i] = 1
		} else {
			freq[i] += 1
		}
	}

	values := make([]int, 0, len(freq))
	for _, v := range freq {
		values = append(values, v)
	}

	slices.Sort(values)

	if (joker) {
		count := freq['J']
		fmt.Println(count)
		if count > 0 && count < 5{
			index := -1 
			for p, v := range values {
				if (v == count) {
					index = p
					break
				}
			}
			if (index == len(values) - 1) {
				values[len(values) - 2] += count 
				values = values[:index]
			} else {
				values[len(values) - 1] += count 
				values = append(values[:index], values[index + 1:]...)
			}
		}
	}

	if Equal(values, []int {5}) {
		hand.handtype = 1
	} else if (Equal(values, []int {1, 4})) {
		hand.handtype = 2
	} else if (Equal(values, []int {2, 3})) {
		hand.handtype = 3
	} else if (Equal(values, []int {1, 1, 3})) {
		hand.handtype = 4
	} else if (Equal(values, []int {1, 2, 2})) {
		hand.handtype = 5
	} else if (Equal(values, []int {1, 1, 1, 2})) {
		hand.handtype = 6
	} else {
		hand.handtype = 7
	}

	return hand
}

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")

	bids := make([]hand, 0, len(lines))
	for _, line := range lines {
		shit := strings.Split(line, " ")
		bids = append(bids, determineHandType(shit[0], shit[1], false))
	}
	ordering := "AKQJT98765432"
	sort.Slice(bids, func(i, j int) bool {
		hand1, hand2 := bids[i], bids[j]
		if (hand1.handtype != hand2.handtype) {
			return hand1.handtype < hand2.handtype
		} else {
			for i := 0; i < 5; i++ {
				if (strings.Index(ordering, string(hand1.hand[i])) != strings.Index(ordering, string(hand2.hand[i]))) {
					return strings.Index(ordering, string(hand1.hand[i])) < strings.Index(ordering, string(hand2.hand[i]))
				}
			}
		}
		return false
	})
	sum := 0
	for i := 0 ; i < len(bids); i++ {
		sum += bids[i].bid * (len(bids) - i)
	}
	fmt.Println(sum)

	ordering = "AKQT98765432J"
	bids = make([]hand, 0, len(lines))
	for _, line := range lines {
		shit := strings.Split(line, " ")
		bids = append(bids, determineHandType(shit[0], shit[1], true))
	}
	sort.Slice(bids, func(i, j int) bool {
		hand1, hand2 := bids[i], bids[j]
		if (hand1.handtype != hand2.handtype) {
			return hand1.handtype < hand2.handtype
		} else {
			for i := 0; i < 5; i++ {
				if (strings.Index(ordering, string(hand1.hand[i])) != strings.Index(ordering, string(hand2.hand[i]))) {
					return strings.Index(ordering, string(hand1.hand[i])) < strings.Index(ordering, string(hand2.hand[i]))
				}
			}
		}
		return false
	})
	sum = 0
	for i := 0 ; i < len(bids); i++ {
		sum += bids[i].bid * (len(bids) - i)
	}
	fmt.Println(sum)
}