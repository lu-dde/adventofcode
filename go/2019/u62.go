package main

import (
	"fmt"
	"strings"
)

//U62 is main proxy for solve, takes a string channel
func U62(p chan string, s chan string) {

	var rel = map[string][]string{}

	// format: QWE)ERT, build tree with root COM
	for relation := range p {
		parts := strings.Split(relation, ")")
		inner := parts[0]
		outer := parts[1]

		_, ok := rel[inner]
		if ok {
			rel[inner] = append(rel[inner], outer)
		} else {
			rel[inner] = []string{outer}
		}
	}

	depthSum := sumTreeDepth(rel, "COM", 0)

	s <- fmt.Sprintf("Solution: %v", depthSum)
}
