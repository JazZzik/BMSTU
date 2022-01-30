package main

import (
	"fmt"
)

func printAutomata(conditionsTo [][]int, output[][]string, startCond int)  {
	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")
	fmt.Println("dummy [label = \"\", shape = none]")
	for i := 0; i < len(conditionsTo); i++ {
		fmt.Printf("%d [shape = circle]\n", i)
	}
	fmt.Printf("dummy -> %d\n", startCond)
	for i := 0; i < len(conditionsTo); i++ {
		for j := 0; j < len(conditionsTo[i]); j++ {
			fmt.Printf("%d -> %d [label = \"%s(%s)\"]\n", i, conditionsTo[i][j], string(j + 97), output[i][j])
		}
	}
	fmt.Println("}")
}

func main()  {
	var i, j, k int
	var s string
	var conN int
	_, _ = fmt.Scan(&conN)
	var input int
	_, _ = fmt.Scan(&input)
	var startCond int
	_, _ = fmt.Scan(&startCond)
	conditionsTo := make([][]int, conN)
	output := make([][]string, conN)
	for i = 0; i < conN; i++ {
		conditionsTo[i] = make([]int, input)
		for j = 0; j < input; j++ {
			_, _ = fmt.Scan(&k)
			conditionsTo[i][j] = k
		}
	}
	for i = 0; i < conN; i++ {
		output[i] = make([]string, input)
		for j = 0; j < input; j++ {
			_, _ = fmt.Scan(&s)
			output[i][j] = s
		}
	}
	printAutomata(conditionsTo, output, startCond)
}
