package ch5

import (
	"fmt"
	"regexp"
)

var mobileRegexp = regexp.MustCompile(`1\d10`)
var mobileCode = regexp.MustCompile(`[0-9a-zA-Z]{4}`)
var s = regexp.MustCompile(`,`)

func test() {
	fmt.Println(mobileRegexp.MatchString("18335831710"))
	fmt.Println(mobileCode.FindString("--1453-"))
	fmt.Println(mobileCode.Split("---4536+++", 3))
	fmt.Println(s.FindAllString("da,d,d,a,d,d", -1))

}
