package main

import (
	"fmt"
)

var c = 0

type condition struct {
	number int
	canonNumber int
	to []*condition
	use bool
}

func canonize(conditionsTo [][]int, output[][]string, startCond int, conN int)  ([][]int, [][]string, int){
	condoms := make([]condition, conN)
	for i := 0; i < conN; i++ {
		var newCond condition
		newCond.number = i
		newCond.canonNumber = -1
		newCond.use = false
		newCond.to = make([]*condition, 0)
		condoms[i] = newCond
	}
	for i := 0; i < conN; i++ {
		for j := 0; j < len(conditionsTo[0]); j++ {
			condoms[i].to = append(condoms[i].to, &condoms[conditionsTo[i][j]])
		}
	}
	visitCondition(&condoms[startCond])
	newCondoms := make([]condition, 0)
	for _, condom := range condoms {
		if condom.canonNumber == -1 {
			conN--
		} else {
			newCondoms = append(newCondoms, condom)
		}
	}
	canonConditionsTo := make([][]int, len(newCondoms))
	canonOutput := make([][]string, len(newCondoms))
	for _, condom := range newCondoms {
		canonConditionsTo[condom.canonNumber] = make([]int, len(conditionsTo[0]))
		canonOutput[condom.canonNumber] = make([]string, len(conditionsTo[0]))
		for j, condomTo := range condom.to {
			canonConditionsTo[condom.canonNumber][j] = condomTo.canonNumber
			canonOutput[condom.canonNumber][j] = output[condom.number][j]
		}
	}
	return canonConditionsTo, canonOutput, conN
}
func visitCondition(condom *condition)  {
	condom.use = true
	condom.canonNumber = c
	c++
	for _, condomTo := range condom.to {
		if !condomTo.use {
			visitCondition(condomTo)
		}
	}
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
	conditionsTo, output, conN = canonize(conditionsTo, output,startCond, conN)
	fmt.Printf("%d\n", conN)
	fmt.Printf("%d\n", input)
	fmt.Println(0)
	for i = 0; i < conN; i++ {
		for j = 0; j < input; j++ {
			fmt.Printf("%d ", conditionsTo[i][j])
		}
		fmt.Println()
	}
	for i = 0; i < conN; i++ {
		for j = 0; j < input; j++ {
			fmt.Printf("%s ", output[i][j])
		}
		fmt.Println()
	}
}