package main

import (
	"fmt"
	//"github.com/skorobogatov/input"
)


type vertex struct {
	mark int
	marked bool
	edges []int
}

type component struct{
	numNod int
	numVert int
	group int
}

var (
	incident []vertex
	group, maxgroup  component
	groupNum = 0
)

func makeIncident(i int) {
	incident[i].marked = true
	incident[i].mark, group.group = groupNum, groupNum
	group.numNod++
	group.numVert += len(incident[i].edges)
	for _, v := range incident[i].edges{
		if !incident[v].marked { makeIncident(v) }
	}
}

func main() {
	var n, m, v1, v2 int
	var ed [2]int
	_, _ = fmt.Scanf("%d\n%d", &n, &m)
	//input.Scanf("%d\n%d", &n, &m)
	edges := make([][2]int, 0)
	for i := 0; i < n; i++{
		var v vertex
		v.mark = -1
		v.marked = false
		v.edges = make([]int, 0)
		incident = append(incident, v)
	}
	for i := 0; i < m; i++ {
		//input.Scanf("%d %d", &v1, &v2)
		_, _ = fmt.Scanf("%d %d", &v1, &v2)
		ed[0] = v1
		ed[1] = v2
		edges = append(edges, ed)
		incident[v1].edges = append(incident[v1].edges, v2)
		incident[v2].edges = append(incident[v2].edges, v1)
	}

	for i := range incident{
		group.numNod, group.numVert = 0, 0
		if !incident[i].marked { makeIncident(i) }
		if group.numNod > maxgroup.numNod || (group.numNod == maxgroup.numNod && group.numVert > maxgroup.numVert) { maxgroup = group }
		groupNum++
	}
	fmt.Printf("graph {\n")
	for i, v := range incident{
		if maxgroup.group == v.mark{
			fmt.Printf("\t%d [color = red]\n", i)
		} else {
			fmt.Printf("\t%d\n", i)
		}
	}
	for _, v := range edges{
		if maxgroup.group == incident[v[0]].mark{
			fmt.Printf("\t%d -- %d [color = red]\n", v[0], v[1])
		} else {
			fmt.Printf("\t%d -- %d\n", v[0], v[1])
		}
	}
	fmt.Printf("}")
}
