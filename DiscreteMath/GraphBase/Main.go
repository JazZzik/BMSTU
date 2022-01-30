package main

import (
	"fmt"
	"github.com/skorobogatov/input"
)

type vertex struct {
	comp, low, t1 int
	edges []int
	in bool
}

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


var (
	time  = 1
	count = 0
)

func Tarjan(incident *[]vertex) {
	var s stack
	for i := range *incident {
		if (*incident)[i].t1 == 0 {
			VisitVertex_Tarjan(incident, i, &s)
		}
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
		(*incident)[u].comp = count
		for u != v {
			u = s.pop()
			(*incident)[u].comp = count
		}
		count++
	}
}

func main() {
	var m, n, v1, v2 int
	//_, _ = fmt.Scanf("%d\n%d", &n, &m)
	input.Scanf("%d\n%d", &n, &m)
	incident := make([]vertex, 0)
	var v vertex
	for i := 0; i < n; i++ {
		v.edges = make([]int, 0)
		v.comp = -1
		incident = append(incident, v)
	}
	for i := 0; i < m; i++ {
		//_, _ = fmt.Scanf("%d %d", &v1, &v2)
		input.Scanf("%d %d", &v1, &v2)
		incident[v1].edges = append(incident[v1].edges, v2)
	}
	Tarjan(&incident)
	condensation := make([]vertex, 0)
	for i := 0; i < count; i++ {
		v.edges = make([]int, 0)
		v.comp = i
		condensation = append(condensation, v)
	}
	for i := range incident{
		for _, u := range incident[i].edges{
			if incident[i].comp != incident[u].comp{ condensation[incident[u].comp].in = true }
		}
	}
	base := make([]int, 0)
	for _, c := range condensation{
		if !c.in {
			for j, u := range incident{
				if c.comp == u.comp {
					base = append(base, j)
					break
				}
			}
		}
	}
	for _, u := range base { fmt.Printf("%d ", u) }
}