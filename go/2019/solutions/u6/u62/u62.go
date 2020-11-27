package u62

import (
	"fmt"
	"strings"
)

type treeMapNode struct {
	parent string
	leaf   []string
}

type treeMapDistance struct {
	node     string
	distance int
}

func orbitWalkToTop(treeMap map[string]treeMapNode, from string) []treeMapDistance {
	var walk = []treeMapDistance{}
	var distance = 0

	for treeMap[from].parent != "" {
		from = treeMap[from].parent
		walk = append(walk, treeMapDistance{node: from, distance: distance})
		distance++
	}

	return walk
}

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

	var treeMap = map[string]treeMapNode{}

	// format: QWE)ERT, build treeMapNode with root COM
	for relation := range p {
		parts := strings.Split(relation, ")")
		inner := parts[0]
		outer := parts[1]

		innerNode, okInner := treeMap[inner]
		if okInner {
			treeMap[inner] = treeMapNode{parent: innerNode.parent, leaf: append(innerNode.leaf, outer)}
		} else {
			treeMap[inner] = treeMapNode{leaf: []string{outer}}
		}

		orbitNode, okOuter := treeMap[outer]
		if okOuter {
			treeMap[outer] = treeMapNode{parent: inner, leaf: orbitNode.leaf}
		} else {
			treeMap[outer] = treeMapNode{parent: inner}
		}
	}

	you := orbitWalkToTop(treeMap, "YOU")
	san := orbitWalkToTop(treeMap, "SAN")

	/*
		fmt.Println("YOU", you)
		fmt.Println("SAN", san)
	*/

	sanOffset := len(you) - len(san)

	for i := len(you) - 1; i >= 0 && i-sanOffset >= 0; i-- {
		/*
			fmt.Print("YOU ", i, " ", you[i].node)
			fmt.Print(" SAN ", i-sanOffset, " ", san[i-sanOffset].node)
			fmt.Println()
		*/
		if you[i].node != san[i-sanOffset].node {
			orbits := you[i+1].distance + you[i+1-sanOffset].distance
			s <- fmt.Sprintf("Solution: %s -> %s = %d", treeMap["YOU"].parent, treeMap["SAN"].parent, orbits)
			break
		}
	}
}
