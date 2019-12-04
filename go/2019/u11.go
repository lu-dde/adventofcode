package main

import "fmt"

//U11 is main proxy for solve, takes a string channel
func U11(p chan string, s chan string) {
	solve()
}

func solve() {
	fmt.Println("u11")
}
