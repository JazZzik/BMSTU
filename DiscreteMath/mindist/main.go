package main

import (
	"fmt"
	//"github.com/skorobogatov/input"
)

func Abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	var str string
	var	x, y rune
	//str = input.Gets()
	//input.Scanf("%c %c", &x, &y)
	_, _ = fmt.Scanf("%s\n", &str)
	_, _ = fmt.Scanf("%c %c", &x, &y)
	xi, yi, i, dist := -1, -1, 0, 1000000
	for _, s := range str {
		if s == x { xi = i }
		if s == y { yi = i }
		if xi != -1 && yi != -1 { if Abs(yi - xi) < dist { dist = Abs(yi - xi) } }
		i++
	}
	fmt.Println(dist - 1)
}
