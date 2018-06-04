package main

import (
	"gostudy/tools"
)

func main() {
	tools.ExecClock()
}

func twoSum(nums []int, target int) []int {
	s := make([]int, 2)
	for pi, pv := range nums {
		for i, v := range nums {
			if pv+v == target && pi != i {
				s[0], s[1] = pi, i
				return s
			}
		}
	}
	return s
}
