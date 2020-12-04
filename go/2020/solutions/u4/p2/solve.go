package p2

import (
	"fmt"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(input chan string, s chan string) {
	var t = 0

	last := map[string]string{}

	for line := range input {
		if line == "" {
			if check(last) {
				t++
			}
			last = make(map[string]string)
			continue
		}

		for _, kv := range strings.Fields(line) {
			kvs := strings.Split(kv, ":")
			key := kvs[0]
			value := kvs[1]
			last[key] = value
		}
	}

	if check(last) {
		t++
	}

	s <- fmt.Sprintf("Solution: %d", t)
}

func check(m map[string]string) bool {
	l := len(m)
	return l == 8 || (l == 7 && m["cid"] == "")
}
