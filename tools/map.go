package tools

import (
	"fmt"
	"unicode"
)

type Student struct {
	name  string
	sex   byte
	brith int
}

// 学生map
func MyMap() {
	students := make(map[string]Student)
	students["001"] = Student{name: "hanjian", sex: 1, brith: 186522}
	if stu, ok := students["001"]; !ok {
		fmt.Println("该学生不存在")
	} else {
		updateStudent(&stu)
		fmt.Println(stu)
	}
}

func updateStudent(stu *Student) {
	stu.sex = 2
}

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
