package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")

	sum := 0
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum = sum + massToFuelRecursive(mass)
	}

	fmt.Printf("Result is %d\n", sum)
}

func massToFuel(mass int) int {
	return int(math.Floor(float64(mass/3)) - 2)
}

func massToFuelRecursive(mass int) int {
	fuel := massToFuel(mass)

	if fuel <= 0 {
		return 0
	}

	return fuel + massToFuelRecursive(fuel)
}
