package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	opcodeAdd       = 1
	opcodeMultiply  = 2
	opcodeInput     = 3
	opcodeOutput    = 4
	opcodeJumpIf    = 5
	opcodeJumpIfNot = 6
	opcodeLessThan  = 7
	opcodeEqual     = 8
	opcodeEnd       = 99

	modePosition  = 0
	modeImmediate = 1
)

func main() {
	input := strings.Split(os.Args[1], ",")

	data := make([]int, len(input))
	for i, val := range input {
		res, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		data[i] = res
	}

	process(data)
}

func parseOperand(opOrdinal, instructionIndex int, data []int) int {
	mode := data[instructionIndex] / int(math.Pow10(opOrdinal+1)) % 10
	op := data[instructionIndex+opOrdinal]
	if mode == modePosition {
		op = data[op]
	}

	return op
}

func parseInstruction(instruction, index int, data []int) (opcode int, parameters []int) {
	opcode = instruction % 100
	parameters = nil

	switch opcode {
	case opcodeMultiply:
		fallthrough
	case opcodeLessThan:
		fallthrough
	case opcodeEqual:
		fallthrough
	case opcodeAdd:
		op1 := parseOperand(1, index, data)
		op2 := parseOperand(2, index, data)
		op3 := data[index+3]

		parameters = []int{op1, op2, op3}
	case opcodeOutput:
		op := parseOperand(1, index, data)

		parameters = []int{op}
	case opcodeJumpIfNot:
		fallthrough
	case opcodeJumpIf:
		op1 := parseOperand(1, index, data)
		op2 := parseOperand(2, index, data)
		parameters = []int{op1, op2}
	case opcodeInput:
	case opcodeEnd:
	}

	return
}

func process(data []int) {
	i := 0
	for {
		instruction := data[i]
		opcode, params := parseInstruction(instruction, i, data)
		switch opcode {
		case opcodeAdd:
			data[params[2]] = params[0] + params[1]
			i = i + 4
		case opcodeMultiply:
			data[params[2]] = params[0] * params[1]
			i = i + 4
		case opcodeInput:
			var input int
			_, err := fmt.Scanf("%d\n", &input)
			if err != nil {
				panic(err)
			}

			data[data[i+1]] = input
			i = i + 2
		case opcodeOutput:
			fmt.Println(params[0])
			i = i + 2
		case opcodeJumpIf:
			if params[0] != 0 {
				i = params[1]
			} else {
				i = i + 3
			}
		case opcodeJumpIfNot:
			if params[0] == 0 {
				i = params[1]
			} else {
				i = i + 3
			}
		case opcodeLessThan:
			if params[0] < params[1] {
				data[params[2]] = 1
			} else {
				data[params[2]] = 0
			}
			i = i + 4
		case opcodeEqual:
			if params[0] == params[1] {
				data[params[2]] = 1
			} else {
				data[params[2]] = 0
			}
			i = i + 4
		case opcodeEnd:
			return
		}
	}
}
