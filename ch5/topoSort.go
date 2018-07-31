package ch5

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) (order []string) {
	var keys []string

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	visitAll(keys)
	return
}

//匿名函数存放的是函数中的变量的地址
func n() (sum int) {
	for i := 1; i <= 100; i++ {
		i := i
		sum += i
		defer func() { fmt.Println(i) }()
	}
	return sum
}

func sum(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

func sumSlice(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}
