package p1

import (
	"fmt"
	"regexp"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {
	var t = 0

	rules := rules{
		R: map[string]rule{},
	}

	for line := range p {
		if line == "" {
			break
		}

		//fmt.Println(line)

		p := strings.Split(line, ": ")
		ruleID := p[0]
		p2 := p[1]
		var subRules [][]string
		or := strings.Split(p2, " | ")
		for _, q := range or {
			subRules = append(subRules, strings.Fields(q))
		}

		rules.add(ruleID, subRules)
	}

	var b strings.Builder
	//b.Grow(2048)
	b.WriteByte('^')
	rules.buildRegex("0", &b)
	b.WriteByte('$')

	//fmt.Println(b.String())

	check := regexp.MustCompile(b.String())

	for line := range p {
		if check.MatchString(line) {
			t++
			//	fmt.Println(line)
		} else {
			//	fmt.Println(line, "not")
		}
	}

	s <- fmt.Sprintf("Solution: %d", t)
}

type rule struct {
	ID   string
	Type string
	sub  [][]string
}

type rules struct {
	R map[string]rule
}

func (rr *rules) buildRegex(id string, b *strings.Builder) {
	r := rr.R[id]
	if r.Type == "node" {
		b.WriteByte('(')
		for i, sub := range r.sub {
			if i > 0 {
				b.WriteByte('|')
			}
			for _, s := range sub {
				rr.buildRegex(s, b)
			}
		}
		b.WriteByte(')')
	} else {
		b.WriteString(r.sub[0][0])
	}
}

func (rr *rules) add(id string, sub [][]string) {
	r := rule{
		ID:   id,
		sub:  sub,
		Type: "node",
	}

	if sub[0][0][0] == '"' {
		r.sub = [][]string{
			{
				string(sub[0][0][1]),
			},
		}
		r.Type = "leaf"
	}

	rr.R[id] = r
}
