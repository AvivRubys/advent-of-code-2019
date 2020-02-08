package main

import (
	"fmt"
)

const (
	MinValue = 235741
	MaxValue = 706948
	Mask     = 100000
)

func main() {
	count1, count2 := 0, 0

	for i := MinValue; i <= MaxValue; i++ {
		if !isIncreasing(i) {
			continue
		}

		groups := groupByDigit(i)

		for _, grp := range groups {
			if grp >= 2 {
				count1++
				break
			}
		}

		for _, grp := range groups {
			if grp == 2 {
				count2++
				break
			}
		}
	}

	fmt.Println(count1)
	fmt.Println(count2)
}

func isIncreasing(n int) bool {
	prev := n / Mask

	for mask := Mask / 10; mask > 0; mask = mask / 10 {
		curr := n / mask % 10
		if prev > curr {
			return false
		}

		prev = curr
	}

	return true
}

func groupByDigit(n int) []int {
	groups := make([]int, 10)

	for mask := Mask; mask > 0; mask = mask / 10 {
		groups[n/mask%10]++
	}

	return groups
}
