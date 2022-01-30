package main

import (
	"fmt"
	"os"
	"sort"
)

type condition struct {
	number int
	canonNumber int
	to []*condition
	use bool
	treeAnsestor *condition
	treeDepth int
	repair int
}

var c = 0

func printAutomata(conditionsTo [][]int, output[][]string, startCond int) {
	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")
	fmt.Println("\tdummy [label = \"\", shape = none]")
	for i := 0; i < len(conditionsTo); i++ {
		fmt.Printf("\t%d [shape = circle]\n", i)
	}
	fmt.Printf("\tdummy -> %d\n", startCond)
	for i := 0; i < len(conditionsTo); i++ {
		for j := 0; j < len(conditionsTo[i]); j++ {
			fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n", i, conditionsTo[i][j], string(j + 97), output[i][j])
		}
	}
	fmt.Println("}")
}

func canonize(conditionsTo [][]int, output[][]string, startCond int) ([][]int, [][]string){
	conN := len(conditionsTo)
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
	return canonConditionsTo, canonOutput
}

func visitCondition(condom *condition) {
	condom.use = true
	condom.canonNumber = c
	c++
	for _, condomTo := range condom.to {
		if !condomTo.use {
			visitCondition(condomTo)
		}
	}
}

func find(condom condition)  *condition {
	copyCon := &condom
	for copyCon.treeAnsestor != copyCon {
		copyCon = copyCon.treeAnsestor
	}
	return copyCon
}

func union (condomOne, condomTwo condition)  {
	ansest1 := find(condomOne)
	ansest2 := find(condomTwo)
	if ansest1.treeDepth == ansest2.treeDepth {
		ansest1.treeDepth++
		ansest2.treeAnsestor = ansest1
	} else {
		if ansest1.treeDepth > ansest2.treeDepth {
			ansest2.treeAnsestor = ansest1
		} else {
			ansest1.treeAnsestor = ansest2
		}
	}
}

func split1(output[][]string)  ([]*condition, int){
	alphabet := len(output[0])
	condoms := make([]condition, len(output))
	for i := 0; i < len(output); i++ {
		condoms[i].treeAnsestor = &condoms[i]
		condoms[i].number = i
	}
	m := len(condoms)
	PPSh := make([]*condition, m)
	for i := 0; i < len(condoms); i++ {
		for j := i + 1; j < len(condoms); j++ {
			if find(condoms[i]) != find(condoms[j]) {
				eq := true
				for a := 0; a < alphabet; a++ {
					if output[i][a] != output[j][a] {
						eq = false
						break
					}
				}
				if eq {
					union(condoms[i], condoms[j])
					m--
				}
			}
		}
	}

	for i := 0; i < len(condoms); i++ {
		PPSh[i] = find(condoms[i])
	}
	return PPSh, m
}

func split(conditionsTo[][]int, PPSh []*condition)  ([]*condition, int) {
	alphabet := len(conditionsTo[0])
	condoms := make([]condition, len(conditionsTo))
	for i := 0; i < len(conditionsTo); i++ {

		condoms[i].treeAnsestor = &condoms[i]
		condoms[i].number = i

	}
	m := len(condoms)
	for i := 0; i < len(condoms); i++ {
		for j := i + 1; j < len(condoms); j++ {
			if find(condoms[i]) != find(condoms[j]) && PPSh[i] == PPSh[j] {
				eq := true
				for a := 0; a < alphabet; a++ {
					w1 := conditionsTo[i][a]
					w2 := conditionsTo[j][a]
					if PPSh[w1] != PPSh[w2] {
						eq = false
						break
					}
				}
				if eq {
					union(condoms[i], condoms[j])
					m--
				}
			}
		}
	}

	for i := 0; i < len(condoms); i++ {
		PPSh[i] = find(condoms[i])
	}
	return PPSh, m
}

func repairIndexes(PPSh []*condition)  []*condition{
	PPCh := make([]*condition, 0)
	for i := 0; i < len(PPSh); i++ {
		PPCh = append(PPCh, PPSh[i])
	}
	sort.Slice(PPCh, func(i, j int) bool {
		return PPCh[i].number < PPCh[j].number
	})
	r := 0
	for i := 0; i < len(PPCh); i++ {
		PPCh[i].repair = r
		if i != len(PPCh) - 1 {
			if PPCh[i + 1].number != PPCh[i].number {
				r++
			}
		}
	}

	for i := 0; i < len(PPSh); i++ {
		for j := 0; j < len(PPCh); j++ {
			if PPSh[i].number == PPCh[j].number {
				PPSh[i].repair = PPCh[j].repair
			}
		}
	}
	return PPSh
}

func In(n []int, I int)  bool {
	for _, a := range n {
		if a == I {
			return true
		}
	}
	return false
}

func AufenkampHohn(conditionsTo[][]int, output[][]string, startCond int) ([][]int, [][]string) {
	var m1 int
	alphabet := len(conditionsTo[0])
	PPSh, m := split1(output)

	for {
		PPSh, m1 = split(conditionsTo, PPSh)
		if m == m1 {
			break
		}
		m = m1
	}

	PPSh = repairIndexes(PPSh)

	startCond = PPSh[startCond].repair
	newConditionsTo := make([][]int, 0)
	newOutput := make([][]string, 0)
	repeat := make([]int, 0)
	for i := 0; i < len(conditionsTo); i++ {
		condom := PPSh[i]
		if !In(repeat, condom.repair) {
			repeat = append(repeat, condom.repair)
			toRow := make([]int, alphabet)
			outRow := make([]string, alphabet)
			for a := 0; a < alphabet; a++ {
				toRow[a] = PPSh[conditionsTo[i][a]].repair
				outRow[a] = output[i][a]
			}
			newConditionsTo = append(newConditionsTo, toRow)
			newOutput = append(newOutput, outRow)
		}
	}
	newConditionsTo, newOutput = canonize(newConditionsTo, newOutput, startCond)
	return newConditionsTo, newOutput
	//printAutomata(newConditionsTo, newOutput, 0)
}

func equal (conditionsTo[][]int, output[][]string, conditionsTo1[][]int, output1[][]string) {
	if len(conditionsTo) == len(conditionsTo1) && len(conditionsTo[0]) == len(conditionsTo1[0]) {
		for i := 0; i < len(conditionsTo); i++ {
			for j := 0; j < len(conditionsTo[0]); j++ {
				if conditionsTo[i][j] != conditionsTo1[i][j] || output[i][j] != output1[i][j] {
					fmt.Println("NOT EQUAL")
					os.Exit(0)
				}
			}
		}
		fmt.Println("EQUAL")
	} else {
		fmt.Println("NOT EQUAL")
	}
}

func main() {
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
	conditionsTo, output = AufenkampHohn(conditionsTo, output, startCond)
	c = 0
	_, _ = fmt.Scan(&conN)
	_, _ = fmt.Scan(&input)
	_, _ = fmt.Scan(&startCond)
	conditionsTo1 := make([][]int, conN)
	output1 := make([][]string, conN)
	for i = 0; i < conN; i++ {
		conditionsTo1[i] = make([]int, input)
		for j = 0; j < input; j++ {
			_, _ = fmt.Scan(&k)
			conditionsTo1[i][j] = k
		}
	}
	for i = 0; i < conN; i++ {
		output1[i] = make([]string, input)
		for j = 0; j < input; j++ {
			_, _ = fmt.Scan(&s)
			output1[i][j] = s
		}
	}
	conditionsTo1, output1 = AufenkampHohn(conditionsTo1, output1, startCond)
	equal(conditionsTo, output, conditionsTo1, output1)
}
