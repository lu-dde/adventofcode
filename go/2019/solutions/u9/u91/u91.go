package u91

import (
	"fmt"
	"strconv"
	"strings"
)

//Solve is main proxy for solve, takes a string channel
func Solve(p chan string, s chan string) {

	// we only expect one line.
	line, _ := <-p

	code := strings.Split(line, ",")
	var ops = make([]int64, 4096) // TODO

	for index, c := range code {
		i, _ := strconv.ParseInt(c, 10, 64)
		ops[index] = i
	}

	var input = make(chan int64, 1)
	var output = make(chan int64)
	var solution = make(chan int64)

	input <- 1

	machine := Intcode6{
		ops:    ops,
		input:  input,
		output: output,
	}

	go func() {
		machine.Run()
	}()

	go func() {
		var out int64
		for o := range output {
			out = o
		}
		solution <- out
	}()

	s <- fmt.Sprint("Solution: ", <-solution)

}

type mode int64

const (
	position mode = 0
	direct   mode = 1
	relative mode = 2
)

const (
	opEXIT      int64 = 99
	opADD       int64 = 1
	opMUL       int64 = 2
	opLT        int64 = 7
	opEQ        int64 = 8
	opJMPNEZERO int64 = 5
	opJMPEQZERO int64 = 6
	opINPUT     int64 = 3
	opOUTPUT    int64 = 4
	opRELBASE   int64 = 9
)

func getMode(i int64) mode {
	if i == 0 {
		return position
	} else if i == 1 {
		return direct
	} else if i == 2 {
		return relative
	}
	panic("no such mode")
}

type Intcode6 struct {
	ops         []int64
	pos         int64
	rel         int64
	healthcheck int64
	input       chan int64
	output      chan int64
	op          opcode
}

//NewIntcode6 init with needed fields
func NewIntcode6(ops []int64, input chan int64, output chan int64) Intcode6 {
	return Intcode6{
		ops:    ops,
		input:  input,
		output: output,
	}
}

type opcode struct {
	opcode int64
	modes  []mode
}

//Run the incode machine
func (machine *Intcode6) Run() {
	for machine.exec() {
		//fmt.Println(machine.pos, machine.healthcheck)
	}
	close(machine.output)
}

func (machine *Intcode6) exec() bool {
	machine.setOperation()

	//fmt.Println("opcode", machine.op)
	//fmt.Println("pos", machine.pos)
	//fmt.Println("relbase", machine.rel)
	//fmt.Println(machine.ops)

	switch machine.op.opcode {

	case opEXIT:
		return false

	case opADD:
		machine.add()
	case opMUL:
		machine.mul()

	case opLT:
		machine.lt()
	case opEQ:
		machine.eq()

	case opJMPNEZERO:
		machine.jmpnezero()
	case opJMPEQZERO:
		machine.jmpeqzero()

	case opINPUT:
		machine.getInput()
	case opOUTPUT:
		machine.setOutput()

	case opRELBASE:
		machine.setRel()

	default:
		fmt.Println(machine)
		panic("unknown state")
	}

	return true
}

func (machine *Intcode6) setOperation() {
	opcodeCompact := machine.ops[machine.pos]
	op := opcodeCompact % 100
	modes := []mode{
		getMode((opcodeCompact / 100) % 10),
		getMode((opcodeCompact / 1000) % 10),
		getMode((opcodeCompact / 10000) % 10),
	}
	machine.op = opcode{opcode: op, modes: modes}
}

func (machine *Intcode6) read(paramOffset int64) int64 {
	mode := machine.op.modes[paramOffset]
	pos := machine.pos + paramOffset + 1

	var value int64

	//fmt.Println("modes", machine.op.modes, paramOffset)
	//fmt.Println("pos", mode, position, mode == position)
	//fmt.Println("dir", mode, direct, mode == direct)
	//fmt.Println("rel", mode, relative, mode == relative)

	if mode == position {
		source := machine.ops[pos]
		value = machine.ops[source]

	} else if mode == direct {
		value = machine.ops[pos]

	} else if mode == relative {
		source := machine.ops[pos] + machine.rel
		value = machine.ops[source]

	} else {
		panic("unkown read mode")
	}

	return value
}

func (machine *Intcode6) write(paramOffset, value int64) {
	mode := machine.op.modes[paramOffset]
	pos := machine.pos + paramOffset + 1

	//fmt.Println("modes", machine.op.modes, paramOffset)
	//fmt.Println("pos", mode, position, mode == position)
	//fmt.Println("dir", mode, direct, mode == direct)
	//fmt.Println("rel", mode, relative, mode == relative)

	if mode == position {
		source := machine.ops[pos]
		machine.ops[source] = value

	} else if mode == direct {
		panic("mode 1 not allowed in write")

	} else if mode == relative {
		source := machine.ops[pos] + machine.rel
		machine.ops[source] = value

	} else {
		panic("unkown write mode")
	}

}

func (machine *Intcode6) add() {
	value1 := machine.read(0)
	value2 := machine.read(1)
	machine.write(2, value1+value2)
	machine.pos += 4
}

func (machine *Intcode6) mul() {
	value1 := machine.read(0)
	value2 := machine.read(1)
	machine.write(2, value1*value2)
	machine.pos += 4
}

func (machine *Intcode6) lt() {
	value1 := machine.read(0)
	value2 := machine.read(1)
	machine.write(2, bool2int64(value1 < value2))
	machine.pos += 4
}
func (machine *Intcode6) eq() {
	value1 := machine.read(0)
	value2 := machine.read(1)
	machine.write(2, bool2int64(value1 == value2))
	machine.pos += 4
}

func (machine *Intcode6) jmpnezero() {

	value1 := machine.read(0)
	if value1 != 0 {
		machine.pos = machine.read(1)
	} else {
		machine.pos += 3
	}
}
func (machine *Intcode6) jmpeqzero() {
	value1 := machine.read(0)
	if value1 == 0 {
		machine.pos = machine.read(1)
	} else {
		machine.pos += 3
	}
}

func (machine *Intcode6) getInput() {
	machine.write(0, <-machine.input)
	machine.pos += 2
}
func (machine *Intcode6) setOutput() {
	machine.healthcheck = machine.read(0)
	//fmt.Println("healthcheck ", machine.healthcheck)
	machine.output <- machine.healthcheck
	machine.pos += 2
}

func (machine *Intcode6) setRel() {
	machine.rel += machine.read(0)
	machine.pos += 2
}

func bool2int64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
