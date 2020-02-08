package vm

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	opcodeAdd             = 1
	opcodeMultiply        = 2
	opcodeInput           = 3
	opcodeOutput          = 4
	opcodeJumpIf          = 5
	opcodeJumpIfNot       = 6
	opcodeLessThan        = 7
	opcodeEqual           = 8
	opcodeSetRelativeBase = 9
	opcodeEnd             = 99

	modePosition  = 0
	modeImmediate = 1
	modeRelative  = 2
)

type IntCodeVM struct {
	memory       []int64
	pointer      int64
	relativeBase int64
}

type Output struct {
	Type  string
	Value int64
}

func NewVMFromCodeWithSize(code []string, bufSize int) *IntCodeVM {
	data := make([]int64, bufSize)
	for i, val := range code {
		res, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic(err)
		}

		data[i] = res
	}

	return &IntCodeVM{
		memory:       data,
		pointer:      0,
		relativeBase: 0,
	}
}

func getMode(instruction, paramOrdinal int64) int64 {
	return instruction / int64(math.Pow(10, float64(paramOrdinal+1))) % 10
}

func (vm *IntCodeVM) parseOperand(opOrdinal int64) int64 {
	mode := getMode(vm.memory[vm.pointer], opOrdinal)
	op := vm.memory[vm.pointer+opOrdinal]

	switch mode {
	case modeImmediate:
		break
	case modePosition:
		op = vm.memory[op]
	case modeRelative:
		op = vm.memory[op+int64(vm.relativeBase)]
	}

	return op
}

func (vm *IntCodeVM) parsePositionOperand(opOrdinal int64) int64 {
	mode := getMode(vm.memory[vm.pointer], opOrdinal)
	op := vm.memory[vm.pointer+opOrdinal]

	if mode == modeRelative {
		op = op + int64(vm.relativeBase)
	}

	return op
}

func (vm *IntCodeVM) parseInstruction() (opcode int64, parameters []int64) {
	instruction := vm.memory[vm.pointer]
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
		op1 := vm.parseOperand(1)
		op2 := vm.parseOperand(2)
		op3 := vm.parsePositionOperand(3)

		parameters = []int64{op1, op2, op3}
	case opcodeSetRelativeBase:
		fallthrough
	case opcodeOutput:
		op := vm.parseOperand(1)

		parameters = []int64{op}
	case opcodeJumpIfNot:
		fallthrough
	case opcodeJumpIf:
		op1 := vm.parseOperand(1)
		op2 := vm.parseOperand(2)
		parameters = []int64{op1, op2}
	case opcodeInput:
		op := vm.parsePositionOperand(1)

		parameters = []int64{op}
	case opcodeEnd:
	}

	if os.Getenv("DEBUG") != "" {
		fmt.Printf("Instruction %d, Opcode %s, Params %v\n", instruction, opcodeToString(opcode), parameters)
	}

	return
}

func (vm *IntCodeVM) Process(input <-chan int64, output chan<- Output) {
	for {
		opcode, params := vm.parseInstruction()
		switch opcode {
		case opcodeAdd:
			vm.memory[params[2]] = params[0] + params[1]
			vm.pointer = vm.pointer + 4
		case opcodeMultiply:
			vm.memory[params[2]] = params[0] * params[1]
			vm.pointer = vm.pointer + 4
		case opcodeInput:
			output <- Output{Type: "request_input"}
			vm.memory[params[0]] = <-input
			vm.pointer = vm.pointer + 2
		case opcodeOutput:
			output <- Output{Type: "output", Value: params[0]}
			vm.pointer = vm.pointer + 2
		case opcodeJumpIf:
			if params[0] != 0 {
				vm.pointer = params[1]
			} else {
				vm.pointer = vm.pointer + 3
			}
		case opcodeJumpIfNot:
			if params[0] == 0 {
				vm.pointer = params[1]
			} else {
				vm.pointer = vm.pointer + 3
			}
		case opcodeLessThan:
			if params[0] < params[1] {
				vm.memory[params[2]] = 1
			} else {
				vm.memory[params[2]] = 0
			}
			vm.pointer = vm.pointer + 4
		case opcodeEqual:
			if params[0] == params[1] {
				vm.memory[params[2]] = 1
			} else {
				vm.memory[params[2]] = 0
			}
			vm.pointer = vm.pointer + 4
		case opcodeSetRelativeBase:
			vm.relativeBase = vm.relativeBase + params[0]
			vm.pointer = vm.pointer + 2
		case opcodeEnd:
			close(output)
			return
		}
	}
}

func opcodeToString(code int64) string {
	switch code {
	case opcodeAdd:
		return "Add"
	case opcodeMultiply:
		return "Multiply"
	case opcodeInput:
		return "Input"
	case opcodeOutput:
		return "Output"
	case opcodeJumpIf:
		return "JumpIf"
	case opcodeJumpIfNot:
		return "JumpIfNot"
	case opcodeLessThan:
		return "LessThan"
	case opcodeEqual:
		return "Equal"
	case opcodeSetRelativeBase:
		return "SetRelativeBase"
	case opcodeEnd:
		return "End"
	}

	panic("?????")
}
