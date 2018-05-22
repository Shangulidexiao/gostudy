package tools

import "fmt"

func ellipse() {
	cx, cy := 260, 260
	fmt.Println("<svg width='100%' height='100%' version='1.1' xmlns='http://www.w3.org/2000/svg'>")
	for rx, ry := 200, 80; ry > 0; rx, ry = rx-2, ry-1 {
		fmt.Printf("<ellipse cx='%d' cy='%d' rx='%d' ry='%d' "+
			"style='fill: white;stroke:rgb(0,0,100);stroke-width:0.5'/>\n", cx, cy, rx, ry)
	}
	fmt.Println("</svg>")
}
