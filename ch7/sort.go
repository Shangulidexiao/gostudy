package ch7

import (
	"fmt"
	"sort"
)

type Student struct {
	ID   string
	Name string
	age  int
}

var stus = []Student{
	{"00001", "hanjian", 23},
	{"00002", "liting", 34},
}

type byAge []Student

func (s byAge) Len() int {
	return len(s)
}

func (s byAge) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s byAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortStudent() {
	fmt.Println(&stus[0])
	sort.Sort(byAge(stus))
	fmt.Println(&stus[0])

	values := []int{6, 3, 7, 9, 2}
	sort.Ints(values)
	fmt.Println(values)
}
