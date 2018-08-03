//整数集合操作
package set

import (
	"strconv"
	"strings"
)

//整数集合结构体
type IntSet struct {
	words []uint
}

const UINTLEN = 32 << (^uint(0) >> 63)

//向集合中添加一个元素
func (s *IntSet) Add(x int) {
	word, offset := x/UINTLEN, uint(x%UINTLEN)
	if word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << offset
}

//向集合中添加一个元素
func (s *IntSet) Adds(args ...int) {
	for _, x := range args {
		s.Add(x)
	}
}

//从集合中移除一个元素
func (s *IntSet) Remove(x int) {
	word, offset := x/UINTLEN, uint(x%UINTLEN)
	if word <= len(s.words) {
		s.words[word] ^= 1 << offset
	}
}

//判断元素是否在集合中
func (s *IntSet) Has(x int) bool {
	word, offset := x/UINTLEN, uint(x%UINTLEN)
	return word < len(s.words) && s.words[word]&(1<<offset) != 0
}

//集合并集
func (s *IntSet) UnitWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//集合交集
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, _ := range s.words {
		var tword uint
		if i <= len(t.words) {
			tword = t.words[i]
		} else {
			tword = uint(0)
		}
		s.words[i] &= tword
	}
}

//格式化输出结构体
func (s *IntSet) String() string {
	var intSlice []string
	for i, word := range s.words {
		for j := 0; j < UINTLEN; j++ {
			if word&(1<<uint(j)) != 0 {
				intSlice = append(intSlice, strconv.Itoa(i*UINTLEN+j))
			}
		}
	}

	return strings.Join(intSlice, " ")
}

//列出集合中的元素
func (s *IntSet) Elmes() []int {
	var intSlice []int
	for i, word := range s.words {
		for j := 0; j < UINTLEN; j++ {
			if word&(1<<uint(j)) != 0 {
				intSlice = append(intSlice, i*UINTLEN+j)
			}
		}
	}
	return intSlice
}

//获取集合长度
func (s *IntSet) Len() int {
	return len(s.Elmes())
}

//清空集合
func (s *IntSet) Clear() {
	slen := s.Len()
	s.words = make([]uint, slen)
}

// 复制集合
func (s *IntSet) Copy() *IntSet {
	var intset IntSet
	set := make([]uint, len(s.words))
	copy(set, s.words)
	intset.words = set
	return &intset
}
