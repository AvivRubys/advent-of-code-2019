package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	opcodeAdd      = 1
	opcodeMultiply = 2
	opcodeEnd      = 99
)

func main() {
	input := strings.Split("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,19,5,23,2,23,13,27,1,10,27,31,2,31,6,35,1,5,35,39,1,39,10,43,2,9,43,47,1,47,5,51,2,51,9,55,1,13,55,59,1,13,59,63,1,6,63,67,2,13,67,71,1,10,71,75,2,13,75,79,1,5,79,83,2,83,9,87,2,87,13,91,1,91,5,95,2,9,95,99,1,99,5,103,1,2,103,107,1,10,107,0,99,2,14,0,0", ",")

	data := make([]int, len(input))
	for i, val := range input {
		res, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		data[i] = res
	}

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			newData := make([]int, len(data))
			copy(newData, data)
			newData[1] = i
			newData[2] = j

			// fmt.Println(newData)
			process(newData)

			// fmt.Printf("%d, %d ==> %d\n", i, j, newData[0])
			if newData[0] == 19690720 {
				fmt.Println("FOUND", i, j)
			}
		}
	}
}

func process(data []int) {
	for i := 0; i < len(data); i = i + 4 {
		opcode := data[i]
		switch opcode {
		case opcodeAdd:
			op1 := data[data[i+1]]
			op2 := data[data[i+2]]
			// fmt.Printf("Setting %d to %d+%d\n", data[i+3], op1, op2)
			data[data[i+3]] = op1 + op2
			// fmt.Println(data)
		case opcodeMultiply:
			op1 := data[data[i+1]]
			op2 := data[data[i+2]]
			// if data[i+1] <= 4 || data[i+2] <= 4 {
			// fmt.Printf("Setting %d to %d*%d\n", data[i+3], op1, op2)
			// }
			data[data[i+3]] = op1 * op2
			// fmt.Println(data)
		case opcodeEnd:
			return
		}
	}
}
