package main

import "fmt"
type condition struct {
	number int
	symbol string
}
func printAutomata(conditionsTo [][]int, condoms []condition, alphabet[]string)  {
	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")
	for i := 0; i < len(conditionsTo); i++ {
		fmt.Printf("%d [label = \"(%d,%s)\"]\n", i, condoms[i].number, condoms[i].symbol)
	}
	for i := 0; i < len(conditionsTo); i++ {
		for j := 0; j < len(conditionsTo[i]); j++ {
			fmt.Printf("%d -> %d [label = \"%s\"]\n", i, conditionsTo[i][j], alphabet[j])
		}
	}
	fmt.Println("}")
}
func findConditionIndex(condoms []condition, condom condition) int {
	for i := 0; i < len(condoms); i++ {
		if condom == condoms[i] {
			return i
		}
	}
	return -1
}
func Moorisation(conditionsTo [][]int, output [][]string)  ([][]int, []condition){
	condoms := make([]condition, 0)
	for i := 0; i < len(conditionsTo); i++ {
		for j := 0; j < len(conditionsTo[0]); j++{
			var newCond condition
			newCond.number = conditionsTo[i][j]
			newCond.symbol = output[i][j]
			if findConditionIndex(condoms, newCond) == -1 {
				condoms = append(condoms, newCond)
			}
		}
	}
	conditionsMoore := make([][]int, len(condoms))
	for i := 0; i < len(condoms); i++ {
		conditionsMoore[i] = make([]int, len(conditionsTo[0]))
		for j := 0; j < len(conditionsTo[0]); j++ {
			var helpCondom condition
			helpCondom.number = conditionsTo[condoms[i].number][j]
			helpCondom.symbol = output[condoms[i].number][j]
			conditionsMoore[i][j] = findConditionIndex(condoms, helpCondom)
		}
	}
	return conditionsMoore, condoms
}
func main()  {
	var i, j, k int
	var s string
	var input int
	_, _ = fmt.Scan(&input)
	alphabet := make([]string, input)
	for i = 0; i < input; i++ {
		_, _ = fmt.Scan(&s)
		alphabet[i] = s
	}
	_, _ = fmt.Scan(&k)
	for i = 0; i < k; i++ {
		_, _ = fmt.Scan(&s)
	}
	var conN int
	_, _ = fmt.Scan(&conN)
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
	conditionsMoore, condoms := Moorisation(conditionsTo, output)
	printAutomata(conditionsMoore, condoms, alphabet)
}
