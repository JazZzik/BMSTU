package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var time = 1
type stack []int

func (s *stack) empty() bool{
	return len(*s) == 0
}

func (s *stack) push(e int) {
	*s = append(*s, e)
}

func (s *stack) pop() int{
	if s.empty() {
		return 0
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}
type vertex struct {
	name int
	edges []int
	ancestor []int
	t1 int
	comp int
	low int
	ves int
	color string // "e" - empty, "r" - red, "b" - blue
	out string
	strName string
	path int
}

func alalyzeString(incident *[]vertex, str string, mapOfVerx *map[string]int, veraCount *int)  {
	var vera string
	var newV vertex
	var vName string
	str = strings.Replace(str, " ", "", -1)
	oneCut := strings.Index(str, "<")
	if oneCut == - 1 {
		vera = str
	} else {
		vera = str[:oneCut]
	}
	lArenCut := strings.Index(vera, "(")
	if lArenCut != -1 {
		rArenCut := strings.Index(vera, ")")
		weight := vera[lArenCut + 1:rArenCut]
		vName = vera[:lArenCut]
		ves, _ := strconv.Atoi(weight)
		(*mapOfVerx)[vName] = *veraCount
		newV.name = *veraCount
		newV.comp = -1
		newV.t1 = 0
		newV.edges = make([]int, 0)
		newV.color = "e"
		newV.ves = ves
		newV.out = vera
		newV.strName = vName
		newV.path = ves
		newV.ancestor = make([]int, 0)
		*incident = append(*incident, newV)
		*veraCount++
	} else {
		vName = vera
	}
	if oneCut == - 1 {
		str = ""
	} else {
		str = str[oneCut + 1:]
	}
	var vera2 string
	vera = vName
	for str != "" {
		oneCut = strings.Index(str, "<")
		if oneCut == - 1 {
			vera2 = str
		} else {
			vera2 = str[:oneCut]
		}
		lArenCut = strings.Index(vera2, "(")
		if lArenCut != -1 {
			rArenCut := strings.Index(vera2, ")")
			weight := vera2[lArenCut + 1:rArenCut]
			vName = vera2[:lArenCut]
			ves, _ := strconv.Atoi(weight)
			(*mapOfVerx)[vName] = *veraCount
			newV.name = *veraCount
			newV.comp = -1
			newV.t1 = 0
			newV.edges = make([]int, 0)
			newV.ancestor = make([]int, 0)
			newV.color = "e"
			newV.ves = ves
			newV.out = vera2
			newV.strName = vName
			newV.path = ves
			*incident = append(*incident, newV)
			if !in(*veraCount, (*incident)[(*mapOfVerx)[vera]].edges) {
				(*incident)[(*mapOfVerx)[vera]].edges = append((*incident)[(*mapOfVerx)[vera]].edges, *veraCount)
			}
			*veraCount++
		} else {
			vName = vera2
			nowEd := (*mapOfVerx)[vera2]
			if !in(nowEd, (*incident)[(*mapOfVerx)[vera]].edges) {
				(*incident)[(*mapOfVerx)[vera]].edges = append((*incident)[(*mapOfVerx)[vera]].edges, nowEd)
			}
		}
		vera = vName
		if oneCut == - 1 {
			str = ""
		} else {
			str = str[oneCut + 1:]
		}
	}
}
func Tarjan(incident *[]vertex) {
	var s stack
	for i := range *incident {
		if (*incident)[i].t1 == 0 {
			VisitVertex_Tarjan(incident, i, &s)
		}
	}
}
func Relax(incident *[]vertex, u, w int)  bool {
	isChanged := (*incident)[u].path < (*incident)[w].path + (*incident)[u].ves
	if isChanged {
		(*incident)[u].path = (*incident)[w].path + (*incident)[u].ves
		(*incident)[u].ancestor = make([]int, 0)
		(*incident)[u].ancestor = append((*incident)[u].ancestor, w)
	}
	if (*incident)[u].path == (*incident)[w].path + (*incident)[u].ves {
		(*incident)[u].ancestor = append((*incident)[u].ancestor, w)
	}
	return isChanged
}
func cpm(incident *[]vertex)  {
	isChanged := false
	(*incident)[0].path = (*incident)[0].ves
	for numV := range *incident {
		if (*incident)[numV].color != "b" {
			for _, edge := range (*incident)[numV].edges {
				if (*incident)[edge].color != "b" {
					if Relax(incident, edge, numV) {
						isChanged = true
					}
				}
			}
		}
	}
	if isChanged {
		cpm(incident)
	}
}
func VisitVertex_Tarjan(incident *[]vertex, v int, s *stack) {
	(*incident)[v].t1, (*incident)[v].low = time, time
	time++
	s.push(v)
	for _, u := range (*incident)[v].edges {
		if (*incident)[u].t1 == 0 { VisitVertex_Tarjan(incident, u, s) }
		if (*incident)[u].comp == -1 && (*incident)[v].low > (*incident)[u].low {
			(*incident)[v].low = (*incident)[u].low
		}
	}
	if (*incident)[v].low == (*incident)[v].t1 {
		u := s.pop()
		if in(u, (*incident)[u].edges) {
			(*incident)[u].color = "b"
		}
		(*incident)[u].comp = 1
		for u != v {
			u = s.pop()
			(*incident)[u].comp = 1
			DFSblue(incident, u)
		}
	}
}
func DFSblue(incident *[]vertex, v int)  {
	(*incident)[v].color = "b"
	for _, edge := range (*incident)[v].edges {
		if (*incident)[edge].color != "b" {
			DFSblue(incident, edge)
		}
	}
}
func DFSred(incident *[]vertex, v int)  {
	(*incident)[v].color = "r"
	for _, edge := range (*incident)[v].ancestor {
		if (*incident)[edge].color != "r" {
			DFSred(incident, edge)
		}
	}
}

func in(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main()  {
	incident := make([]vertex, 0)
	mapOfVerx := make(map[string]int)
	scan := bufio.NewScanner(os.Stdin)
	var str string
	var description = ""
	veraCount := 0
	for scan.Scan() {
		str = scan.Text()
		if str == "" {
			break
		}
		description += str
	}
	for description != "" {
		oneCut := strings.Index(description, ";")
		if oneCut != -1 {
			str = description[:oneCut]
			description = description[oneCut + 1:]
		} else {
			str = description
			description = ""
		}
		alalyzeString(&incident, str, &mapOfVerx, &veraCount)
	}
	Tarjan(&incident)
	cpm(&incident)
	mPath := 0
	for _, v := range incident {
		if v.path > mPath && v.color != "b"{
			mPath = v.path
		}
	}
	for numV := range incident {
		if incident[numV].path == mPath && incident[numV].color != "b"{
			DFSred(&incident, numV)
		}
	}
	fmt.Printf("digraph {\n")
	for _, v := range incident {
		if v.color == "e" {
			fmt.Printf("\t%s [label = \"%s\"]\n", v.strName, v.out)
		}
		if v.color == "b" {
			fmt.Printf("\t%s [label = \"%s\", color = blue]\n", v.strName, v.out)
		}
		if v.color == "r" {
			fmt.Printf("\t%s [label = \"%s\", color = red]\n", v.strName, v.out)
		}
	}
	for _, v := range incident {
		for _, edge := range v.edges {
			if v.color == "b" {
				fmt.Printf("\t%s -> %s [color = blue]\n", v.strName, incident[edge].strName)
			} else if v.color == "r" && incident[edge].color == "r" && in(v.name, incident[edge].ancestor) {
				fmt.Printf("\t%s -> %s [color = red]\n", v.strName, incident[edge].strName)
			} else {
				fmt.Printf("\t%s -> %s\n", v.strName, incident[edge].strName)
			}
		}
	}
	fmt.Println("}")
}
