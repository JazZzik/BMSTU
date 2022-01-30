package main

import (
	"fmt"
	"sort"
)

type stack []*vertex

func (s *stack) empty() bool{
	return len(*s) == 0
}

func (s *stack) push(e *vertex) {
	*s = append(*s, e)
}

func (s *stack) pop() *vertex{
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

type vertex struct {
	dead bool
	command string
	time int
	operand int
	parent *vertex
	ancestor *vertex
	dom *vertex
	sdom *vertex
	label *vertex
	to []*vertex
	from []*vertex
	bucket []*vertex
}

var time = 1

func FindMin(v *vertex)  *vertex {
	var min *vertex
	if v.ancestor == nil {
		min = v
	} else {
		var st  stack = make([]*vertex, 0)
		var u = v
		for u.ancestor.ancestor != nil {
			st.push(u)
			u = u.ancestor
		}
		for !st.empty() {
			v = st.pop()
			if v.ancestor.label.sdom.time < v.label.sdom.time { v.label = v.ancestor.label }
			v.ancestor = u.ancestor
		}
		min = v.label
	}
	return min
}

func Dominators(G []*vertex){
	sort.Slice(G, func(i, j int) bool { return G[i].time > G[j].time })
	for _, w := range G {
		if w.time == 1 { continue }
		for _, v := range w.from {
			u := FindMin(v)
			if u.sdom.time < w.sdom.time { w.sdom = u.sdom }
		}
		w.ancestor = w.parent
		w.sdom.bucket = append(w.sdom.bucket, w)
		for _, v := range w.parent.bucket {
			u := FindMin(v)
			if u.sdom == v.sdom {
				v.dom = v.sdom
			} else {
				v.dom = u
			}
		}
		w.parent.bucket = nil
	}
	for _, w := range G {
		if w.time == 1 {
			continue
		}
		if w.dom != w.sdom {
			w.dom = w.dom.dom
		}
	}
	G[len(G) - 1].dom = nil
}

func DFS(v *vertex)  {
	v.dead = false
	v.time = time
	time++
	for e := range v.to {
		if v.to[e].dead {
			v.to[e].parent = v
			DFS(v.to[e])
		}
	}
}

func main()  {
	var n, v1, v2, loops int
	var str string
	_, _ = fmt.Scan(&n)
	incident := make([]*vertex, 0)
	mapName := make(map[int]int)
	for i := 0; i < n; i++ {
		var v vertex
		v.dead = true
		v.to = make([]*vertex, 0)
		v.from = make([]*vertex, 0)
		v.bucket = make([]*vertex, 0)
		v.sdom, v.label = &v, &v
		incident = append(incident, &v)
	}
	for i := range incident {
		_, _ = fmt.Scan(&v1, &str)
		incident[i].command = str
		if str != "ACTION" {
			_, _ = fmt.Scan(&v2)
			incident[i].operand = v2
		}
		mapName[v1] = i
	}
	for i, v := range incident {
		if v.command == "ACTION" {
			if i != n-1 {
				v.to = append(v.to, incident[i+1])
				incident[i+1].from = append(incident[i+1].from, incident[i])
			}
		} else if v.command == "JUMP" {
			v2 = mapName[v.operand]
			v.to = append(v.to, incident[v2])
			incident[v2].from = append(incident[v2].from, incident[i])
		} else {
			v2 = mapName[v.operand]
			v.to = append(v.to, incident[v2])
			incident[v2].from = append(incident[v2].from, incident[i])
			if i < n - 1 {
				v.to = append(v.to, incident[i + 1])
				incident[i + 1].from = append(incident[i + 1].from, incident[i])
			}
		}
	}
	DFS(incident[0])
	for i := 0; i < len(incident); i++ {
		if incident[i].dead {
			incident[i] = incident[len(incident) - 1]
			incident[len(incident) - 1] = nil
			incident = incident[:len(incident) - 1]
			i--
		} else {
			for j := 0; j < len(incident[i].from); j++ {
				if incident[i].from[j].dead {
					incident[i].from[j] = incident[i].from[len(incident[i].from) - 1]
					incident[i].from = incident[i].from[:len(incident[i].from) - 1]
					j--
				}
			}
		}
	}
	Dominators(incident)
	for _, v := range incident {
		for _, u := range v.from {
			for u != v && u!= nil { u = u.dom }
			if u == v {
				loops++
				break
			}
		}
	}
	fmt.Printf("%d", loops)
}