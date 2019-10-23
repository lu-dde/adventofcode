package main

import (
	"fmt"
)

type distancePair struct {
	word     []rune
	other    []rune
	distance int
}

func distance(word []rune, other []rune) int {
	fault := 0
	for i := 0; i < 26; i++ {
		if word[i] != other[i] {
			fault++
		}
	}
	return fault
}

func solve(p chan string, s chan string) {

	words := [250][]rune{{}}
	index := 0
	for {
		line, ok := <-p
		if ok {

			words[index] = []rune(line)
		} else {
			break
		}
		index++
	}

	dp := distancePair{[]rune{}, []rune{}, 30}

	for start, w := range words {
		for oi, o := range words {
			if oi <= start {
				continue
			}

			distance := distance(w, o)
			if distance < dp.distance {
				dp = distancePair{w, o, distance}
			}
		}
	}

	s <- fmt.Sprintf("Solution: %d\n%v\n%v", dp.distance, string(dp.word), string(dp.other))
	return
}
