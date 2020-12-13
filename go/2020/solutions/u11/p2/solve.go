package p2

import (
	"fmt"
)

const width = 100

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	seats := make(map[int]int8, width*width)

	row := 0
	for line := range p {
		for i := 0; i < len(line); i++ {
			if line[i] == 'L' {
				seats[(row*100)+i] = 1
			} else {
				seats[(row*100)+i] = 0
			}
		}
		row++
	}

	/*
		printMap(seats)
		fmt.Println()
	*/

	times := 10000
	for {

		changes := 0

		next := make(map[int]int8, width*width)

		for i, s := range seats {
			if s == 2 && neighbours(i, seats) > 4 {
				next[i] = 1
				changes++
			} else if s == 1 && neighbours(i, seats) == 0 {
				next[i] = 2
				changes++
			} else {
				next[i] = s
			}
		}

		seats = nil
		seats = next

		/*
			fmt.Println(changes)
			printMap(seats)
			fmt.Println()
		*/

		if changes == 0 {
			break
		}
		if times == 2 {
			break
		}
		times++
	}

	occupied := 0
	for _, occ := range seats {
		if occ == 2 {
			occupied++
		}
	}

	t = occupied

	s <- fmt.Sprintf("Solution: %d", t)
}

func neighbours(p int, seats map[int]int8) int {
	l := []int{
		-101, -100, -99,
		-1, +1,
		+99, +100, +101,
	}

	neighbours := 0

	for _, direction := range l {
		cur := p
	outer:
		for {
			cur += direction
			spot, ok := seats[cur]
			if !ok {
				break outer
			}
			switch spot {
			case 2:
				neighbours++
				break outer
			case 1:
				break outer
			}
		}

	}
	return neighbours
}

func printMap(seats map[int]int8) {
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			s, seat := seats[i*100+j]
			if !seat {
				fmt.Print("!")
			} else {
				switch s {
				case 2:
					fmt.Print("#")
				case 1:
					fmt.Print("L")
				case 0:
					fmt.Print(".")
				default:
					fmt.Print("ö")

				}
			}
		}
		fmt.Println()
	}
}
