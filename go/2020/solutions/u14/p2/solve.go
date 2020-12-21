package p2

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t uint64 = 0

	store := make(map[uint64]uint64, 0)

	var m mask

	for line := range p {
		l := strings.Fields(line)

		if l[0] == "mask" {
			m.update(l[2])
			//fmt.Println("M ", l[2])
		} else {
			var index uint64 = 0
			fmt.Sscanf(l[0], "mem[%d]", &index)
			s, _ := strconv.ParseUint(l[2], 10, 35)
			for _, a := range m.mask(index) {
				store[a] = s
			}
		}
	}

	for _, v := range store {
		t += v
	}

	s <- fmt.Sprintf("Solution: %d", t)
}

type mask struct {
	nand uint64
	gen  uint64
	or   uint64
}

func (m *mask) mask(n uint64) []uint64 {
	/*
		fmt.Printf("IN %036b (dec %d)\n", n, n)
		fmt.Printf("G  %036b %d\n", m.gen, bits.OnesCount64(m.gen))
		fmt.Printf("&  %036b\n", m.nand)
		fmt.Printf("|  %036b\n", m.or)
		fmt.Printf("B  %036b\n", n&m.nand|m.or)
	*/
	return m.perm(n)
}
func (m *mask) perm(n uint64) []uint64 {
	base := n&m.nand | m.or
	perms := m.gen
	parts := []uint64{}
	var power uint64 = 1
	for i := 0; i < 36; i++ {
		if perms%2 == 1 {
			parts = append(parts, power)
		}
		perms >>= 1
		power <<= 1
	}

	combinations := 1 << len(parts)

	res := make([]uint64, combinations)

	for i := 0; i < combinations; i++ {
		b := i
		index := 0
		var current uint64 = 0

		for {
			if b == 0 {
				break
			}
			if b&1 == 1 {
				current += parts[index]
			}
			b >>= 1
			index++
		}
		res[i] = base + current
		//fmt.Printf("=  %036b\n", res[i])
	}
	return res
}

func (m *mask) update(l string) {

	var nand, gen, or uint64 = 0, 0, 0

	for _, x := range l {
		gen <<= 1
		or <<= 1
		nand <<= 1
		switch x {
		case 'X':
			gen++
		case '1':
			or++
		case '0':
			nand++
		}
	}

	m.gen = gen
	m.nand = nand
	m.or = or
}
