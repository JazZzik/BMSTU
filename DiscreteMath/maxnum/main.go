package main

import (
	"fmt"
	"strconv"
)

func main() {
	var(k, a1, b1 int
		a, b string
	)
	fmt.Scan(&k)
	var array = make([]string, k)
	for i := range array { fmt.Scan(&array[i]) }
	for i := 0; i < k - 1; i++ {
		a = array[i] + array[i + 1]
		b = array[i + 1] + array[i]
		a1, _ = strconv.Atoi(a)
		b1, _ = strconv.Atoi(b)
		if b1 > a1 { array[i], array[i + 1] = array[i + 1], array[i] }
		for j := i; j > 0; j-- {
			a = array[j - 1] + array[j]
			b = array[j] + array[j - 1]
			a1, _ = strconv.Atoi(a)
			b1, _ = strconv.Atoi(b)
			if b1 > a1 { array[j - 1], array[j] = array[j], array[j - 1] }
		}
	}
	for i := range array { fmt.Print(array[i]) }
}
