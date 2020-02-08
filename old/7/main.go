package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

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

	possiblePhaseSettings := permutations([]int{5, 6, 7, 8, 9})
	maxOutputSignal := -1
	var maxPhaseSettings []int

	for _, phaseSettings := range possiblePhaseSettings {
		channels := []chan int{make(chan int), make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)}
		wg := sync.WaitGroup{}
		for i := 0; i < len(phaseSettings); i++ {
			wg.Add(1)
			go process(&wg, data, channels[i], channels[i+1])
		}

		go func() {
			wg.Add(1)
			defer wg.Done()
			channels[0] <- 0
			channels[0] <- phaseSettings[0]
			channels[1] <- phaseSettings[1]
			channels[2] <- phaseSettings[2]
			channels[3] <- phaseSettings[3]
			channels[4] <- phaseSettings[4]
		}()

		resultSignal := -999999999

		go func() {
			wg.Add(1)
			defer wg.Done()

			for output := range channels[5] {
				resultSignal = output
				channels[0] <- output
			}
		}()

		wg.Wait()

		if resultSignal > maxOutputSignal {
			maxOutputSignal = resultSignal
			maxPhaseSettings = phaseSettings
		}
	}

	fmt.Println("Best output signal is", maxOutputSignal, "using", maxPhaseSettings)
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

func process(wg *sync.WaitGroup, data []int, input <-chan int, output chan<- int) {
	inputIdx := 0
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
			data[data[i+1]] = <-input
			inputIdx++
			i = i + 2
		case opcodeOutput:
			output <- params[0]
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
			wg.Done()
			close(output)
			break
		}
	}
}
