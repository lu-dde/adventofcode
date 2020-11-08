package intcode

import (
	"fmt"
)

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

//V6 machine struct
type V6 struct {
	ops         []int64
	pos         int64
	rel         int64
	healthcheck int64
	input       chan int64
	output      chan int64
	op          opcode
}

//New intcode V6 init with needed fields
func New(ops []int64, input chan int64, output chan int64) V6 {
	return V6{
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
func (machine *V6) Run() {
	for machine.exec() {
	}
	close(machine.output)
}

func (machine *V6) exec() bool {
	machine.setOperation()

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

func (machine *V6) setOperation() {
	opcodeCompact := machine.ops[machine.pos]
	op := opcodeCompact % 100
	modes := []mode{
		getMode((opcodeCompact / 100) % 10),
		getMode((opcodeCompact / 1000) % 10),
		getMode((opcodeCompact / 10000) % 10),
	}
	machine.op = opcode{opcode: op, modes: modes}
}

func (machine *V6) read(paramOffset int64) int64 {
	mode := machine.op.modes[paramOffset]
	pos := machine.pos + paramOffset + 1

	var value int64

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

func (machine *V6) write(paramOffset, value int64) {
	mode := machine.op.modes[paramOffset]
	pos := machine.pos + paramOffset + 1

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

func (machine *V6) add() {
	value1 := machine.read(0)
	value2 := machine.read(1)
	machine.write(2, value1+value2)
	machine.pos += 4
}

func (machine *V6) mul() {
	value1 := machine.read(0)
	value2 := machine.read(1)
	machine.write(2, value1*value2)
	machine.pos += 4
}

func (machine *V6) lt() {
	value1 := machine.read(0)
	value2 := machine.read(1)
	machine.write(2, bool2int64(value1 < value2))
	machine.pos += 4
}
func (machine *V6) eq() {
	value1 := machine.read(0)
	value2 := machine.read(1)
	machine.write(2, bool2int64(value1 == value2))
	machine.pos += 4
}

func (machine *V6) jmpnezero() {

	value1 := machine.read(0)
	if value1 != 0 {
		machine.pos = machine.read(1)
	} else {
		machine.pos += 3
	}
}
func (machine *V6) jmpeqzero() {
	value1 := machine.read(0)
	if value1 == 0 {
		machine.pos = machine.read(1)
	} else {
		machine.pos += 3
	}
}

func (machine *V6) getInput() {
	machine.write(0, <-machine.input)
	machine.pos += 2
}
func (machine *V6) setOutput() {
	machine.healthcheck = machine.read(0)
	//fmt.Println("healthcheck ", machine.healthcheck)
	machine.output <- machine.healthcheck
	machine.pos += 2
}

func (machine *V6) setRel() {
	machine.rel += machine.read(0)
	machine.pos += 2
}

func bool2int64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
