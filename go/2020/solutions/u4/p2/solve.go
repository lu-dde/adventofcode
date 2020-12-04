package p2

import (
	"fmt"
	"regexp"
	"strconv"
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
	checks := 0
	if yearCheck(m["byr"], 1920, 2002) {
		checks++
	}
	if yearCheck(m["iyr"], 2010, 2020) {
		checks++
	}
	if yearCheck(m["eyr"], 2020, 2030) {
		checks++
	}

	if hgt := m["hgt"]; hgt != "" {
		l := len(hgt)
		hgtInt := hgt[0 : l-2]
		hgtUnit := hgt[l-2:]

		if hgtUnit == "in" {
			if yearCheck(hgtInt, 59, 76) {
				checks++
			}
		} else if hgtUnit == "cm" {
			if yearCheck(hgtInt, 150, 193) {
				checks++
			}
		}
	}

	if patternCheck(m["hcl"], `^#[a-f0-9]{6}$`) {
		checks++
	}
	if patternCheck(m["ecl"], `^(amb|blu|brn|gry|grn|hzl|oth)$`) {
		checks++
	}
	if patternCheck(m["pid"], `^[0-9]{9}$`) {
		checks++
	}

	//fmt.Println(m, checks)
	return checks == 7
}

func yearCheck(s string, q, p int) bool {
	i, _ := strconv.Atoi(s)
	return i >= q && i <= p
}
func patternCheck(s string, p string) bool {
	matched, _ := regexp.MatchString(p, s)
	return matched
}
