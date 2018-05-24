package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Point struct {
	X string `圆点`
	Y string `圆点`
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

type NWheel struct {
	Circle Circle
	Spokes int
}

func MyStruct() {
	points := []Point{{X: "1", Y: "1"}, {X: "2", Y: "2"}}
	circles := Circle{Point{"1", "2"}, 3}
	wheels := Wheel{Circle{Point{"1", "2"}, 3}, 20}
	fmt.Println(points, circles, wheels)

	data, err := json.Marshal(points)
	if err != nil {
		log.Fatalf("%s", err)
	}
	p := []Point{}
	fmt.Println(json.Unmarshal(data, &p))
	fmt.Println(p)

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(resp.StatusCode)
	resp.Body.Close()
}
