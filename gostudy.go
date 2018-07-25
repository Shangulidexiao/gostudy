package main

import "gostudy/ch4"

// go 语言学习

func main() {
	ch4.ExecIssuse()
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
