package p2

import (
	"fmt"
)

const width = 100

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	seats := make(map[int]bool, width*width)

	row := 0
	for line := range p {
		for i := 0; i < len(line); i++ {
			if line[i] == 'L' {
				seats[(row*100)+i] = false
			}
		}
		row++
	}

	for {

		changes := 0

		next := make(map[int]bool, 100)

		for i, s := range seats {
			n := neighbours(i, seats)
			if s && n >= 4 {
				next[i] = false
				changes++
			} else if !s && n == 0 {
				next[i] = true
				changes++
			} else {
				next[i] = s
			}
		}

		seats = nil
		seats = next

		if changes == 0 {
			break
		}
	}

	occupied := 0
	for _, occ := range seats {
		if occ {
			occupied++
		}
	}

	t = occupied

	s <- fmt.Sprintf("Solution: %d", t)
}

func neighbours(p int, seats map[int]bool) (neighbours int) {
	l := []int{
		p - 101, p - 100, p - 99,
		p - 1, p + 1,
		p + 99, p + 100, p + 101,
	}

	for _, n := range l {
		if seats[n] {
			neighbours++
		}
	}
	return
}

func printMap(seats map[int]bool) {
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			s, seat := seats[i*100+j]
			if !seat {
				fmt.Print(".")
			} else {
				if s {
					fmt.Print("#")
				} else {
					fmt.Print("L")
				}
			}
		}
		fmt.Println()
	}
}
