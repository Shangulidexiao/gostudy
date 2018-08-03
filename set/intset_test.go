package set

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var s IntSet
	s.Add(1)
	if s.words[0] == 0 {
		t.Error("没有添加进集合中", s.words)
	}

	if s.words[0] != uint(1)<<1 {
		t.Error("添加到集合中的值错误，错误值：", s.words[0], uint64(1)<<1)
	}
}

func TestHas(t *testing.T) {
	addeds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	set := map[bool][]int{
		true:  addeds,
		false: []int{11, 12, 13, 14, 15, 100000},
	}
	var s IntSet
	s.Adds(addeds...)

	for status, xs := range set {
		for _, x := range xs {
			if status != s.Has(x) {
				t.Errorf("\n%d in %s \nexcept:%t \nresult:%t", x, s.String(), status, s.Has(x))
			}
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	var s IntSet
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func ExampleAdd() {
	var s IntSet
	s.Add(1)
}
