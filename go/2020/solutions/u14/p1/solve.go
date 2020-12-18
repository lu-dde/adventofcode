package p1

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t uint64 = 0

	store := make(map[int]uint64, 0)

	var m mask

	for line := range p {
		l := strings.Fields(line)

		if l[0] == "mask" {
			m.update(l[2])
		} else {
			index := 0
			fmt.Sscanf(l[0], "mem[%d]", &index)
			s, _ := strconv.ParseUint(l[2], 10, 35)
			store[index] = m.mask(s)
		}
	}

	for _, v := range store {
		t += v
	}

	s <- fmt.Sprintf("Solution: %d", t)
}

type mask struct {
	and uint64
	or  uint64
}

func (m *mask) mask(n uint64) uint64 {
	/*
		fmt.Printf("IN %036b (dec %d)\n", n, n)
		fmt.Printf("&  %036b\n", m.and)
		fmt.Printf("|  %036b\n", m.or)
		fmt.Printf("=  %036b\n", (n&m.and)|m.or)
		fmt.Println()
	*/
	return n&m.and | m.or
}

func (m *mask) update(l string) {

	var and, or uint64 = 0, 0

	for _, x := range l {
		and <<= 1
		or <<= 1
		switch x {
		case 'X':
			and++
		case '1':
			or++
		}
	}

	m.and = and
	m.or = or
}
