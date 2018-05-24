package tools

import "fmt"

type Student struct {
	name  string
	sex   byte
	brith int
}

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