package main

import "gostudy/splider"

// go 语言学习

func main() {
	splider.Exec()
}

// 计算两个数的和
func TwoSum(nums []int, target int) []int {
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
