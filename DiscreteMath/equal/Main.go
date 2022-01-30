package main

import (
	"fmt"
)

type condition struct {
	i   int
	parent *condition
	depth  int
	use   bool
}

type MealyMachine struct {
	n, m, q0 int
	Q []*condition
	conditionsTo [][]int
	output [][]string
}

type Vertex struct {
	oldName, name int
}

var c = 0

func (q *condition) Find() *condition {
	if q == q.parent { return q }
	return q.parent.Find()
}

func (q *condition) Union(x *condition) {
	if q.depth > x.depth {
		x.parent = q
	} else {
		q.parent = x
		if q.depth == x.depth {
			x.depth++
		}
	}
}

func DFS(G [][]*Vertex, v *Vertex) []int {
	order := make([]int, len(G))
	VisitVertex(G, v, order)
	order = order[0:c]
	return order
}

func VisitVertex(G [][]*Vertex, v *Vertex, order []int) {
	v.name = c
	order[c] = v.oldName
	c++
	for _, u := range G[v.oldName] {
		if u.name == -1 {
			VisitVertex(G, u, order)
		}
	}
}

func (MM *MealyMachine) Canon() *MealyMachine {
	var i, j int
	V := make([]*Vertex, MM.n)
	G := make([][]*Vertex, MM.n)
	for i = 0; i < MM.n; i++ {
		var v Vertex
		v.oldName = i
		v.name = -1
		V[i] = &v
	}
	for i = 0; i < MM.n; i++ {
		for j = 0; j < MM.m; j++ {
			G[i] = append(G[i], V[MM.conditionsTo[i][j]])
		}
	}
	order := DFS(G, V[MM.q0])
	nNew := c
	mNew := MM.m
	Qc := make([]*condition, nNew)
	eta := make([][]int, nNew)
	fi := make([][]string, nNew)
	index := 0
	var cond condition
	for _, i = range order {
		if V[i].name != -1 {
			eta[index] = make([]int, mNew)
			fi[index] = make([]string, mNew)
			cond.i = index
			cond.parent = Qc[index]
			cond.depth = 0
			cond.use = true
			Qc[index] = &cond
			for j = 0; j < MM.m; j++ {
				eta[index][j] = V[MM.conditionsTo[i][j]].name
				fi[index][j] = MM.output[i][j]
			}
			index++
		}
	}
	MM.n = nNew
	MM.m = mNew
	MM.q0 = 0
	MM.Q = Qc
	MM.conditionsTo = eta
	MM.output = fi
	return MM
}

func (MM *MealyMachine) Split1() (int, []*condition) {
	var x int
	var eq bool
	m := MM.n
	for _, q := range MM.Q {
		q.parent = q
		q.depth = 0
	}
	for _, q1 := range MM.Q {
		for _, q2 := range MM.Q {
			if q1.Find() != q2.Find() {
				eq = true
				for x = 0; x < MM.m; x++ {
					if MM.output[q1.i][x] != MM.output[q2.i][x] {
						eq = false
						break
					}
				}
				if eq {
					q1.Union(q2)
					m--
				}
			}
		}
	}
	pi := make([]*condition, MM.n)
	for _, q := range MM.Q { pi[q.i] = q.Find() }
	return m, pi
}

func (MM *MealyMachine) Split(pi []*condition) int {
	m := MM.n
	for _, q := range MM.Q {
		q.parent = q
		q.depth = 0
	}
	var x, w1, w2 int
	var eq bool
	for _, q1 := range MM.Q {
		for _, q2 := range MM.Q {
			if pi[q1.i] == pi[q2.i] && q1.Find() != q2.Find() {
				eq = true
				for x = 0; x < MM.m; x++ {
					w1 = MM.conditionsTo[q1.i][x]
					w2 = MM.conditionsTo[q2.i][x]
					if pi[w1] != pi[w2] {
						eq = false
						break
					}
				}
				if eq {
					q1.Union(q2)
					m--
				}
			}
		}
	}
	for _, q := range MM.Q { pi[q.i] = q.Find() }
	return m
}

func (MM *MealyMachine) AufenkampHohn() *MealyMachine {
	m, pi := MM.Split1()
	var x, mc, q0c, nc int
	var qc *condition
	var cond condition
	for {
		mc = MM.Split(pi)
		if m == mc { break }
		m = mc
	}
	Qc := make([]*condition, MM.n)
	eta := make([][]int, 0, MM.n)
	fi := make([][]string, 0, MM.n)
	nc = 0
	balanse := make([]int, MM.n)
	for _, q := range MM.Q {
		qc = pi[q.i]
		if qc.use {
			balanse[qc.i] = nc
			qc.use = false
			cond.i = nc
			cond.parent = Qc[nc]
			cond.depth = 0
			cond.use = true
			Qc[nc] = &cond
			eta = append(eta, make([]int, MM.m))
			fi = append(fi, make([]string, MM.m))
			for x = 0; x < MM.m; x++ {
				eta[Qc[nc].i][x] = pi[MM.conditionsTo[q.i][x]].i
				fi[Qc[nc].i][x] = MM.output[q.i][x]
			}
			nc++
		}
	}
	q0c = balanse[pi[MM.q0].i]
	for i := 0; i < nc; i++ {
		for x = 0; x < MM.m; x++ {
			eta[i][x] = balanse[eta[i][x]]
		}
	}
	MM.n = nc
	MM.q0 = q0c
	MM.Q = Qc[0:nc]
	MM.conditionsTo = eta
	MM.output = fi
	return MM
}

func  PrintAutomata(MM MealyMachine) {
	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")
	fmt.Println("\tdummy [label = \"\", shape = none]")
	for i := 0; i < MM.n; i++ {
		fmt.Printf("\t%d [shape = circle]\n", i)
	}
	fmt.Printf("\tdummy -> %d\n", MM.q0)
	for i := 0; i < MM.n; i++ {
		for j := 0; j < MM.m; j++ {
			fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n", i, MM.conditionsTo[i][j], string(j + 97), MM.output[i][j])
		}
	}
	fmt.Println("}")
}

func main() {
	var n, m, q0, i, j int
	var MM MealyMachine
	_, _ = fmt.Scan(&n, &m, &q0)
	Q := make([]*condition, n)
	conditionsTo := make([][]int, n)
	output := make([][]string, n)
	for i = 0; i < n; i++ {
		Q[i] = &condition{
			i:   i,
			parent: Q[i],
			depth:  0,
			use:   true,
		}
		conditionsTo[i] = make([]int, m)
		for j = 0; j < m; j++ {
			_, _ = fmt.Scan(&conditionsTo[i][j])
		}
	}
	for i = 0; i < n; i++ {
		output[i] = make([]string, m)
		for j = 0; j < m; j++ {
			_, _ = fmt.Scan(&output[i][j])
		}
	}
	MM.n = n
	MM.m = m
	MM.q0 = q0
	MM.Q = Q
	MM.conditionsTo = conditionsTo
	MM.output = output
	PrintAutomata(*MM.AufenkampHohn().Canon())
}