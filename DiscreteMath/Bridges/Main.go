package main

import (
	"fmt"
	"github.com/skorobogatov/input"
)

type queue struct {
	data []*vertex
	cap,count, head, tail int
}

func (q *queue) empty() bool{
	return (*q).count == 0
}

func (q *queue) init(n int)  {
	(*q).data = make([]*vertex, n)
	(*q).cap = n
	(*q).count = 0
	(*q).head = 0
	(*q).tail = 0
}

func (q *queue) enqueue(x *vertex)  {
	(*q).data[(*q).tail] = x
	(*q).tail++
	if (*q).tail ==  (*q).cap { (*q).tail = 0 }
	(*q).count++
}

func (q *queue) dequeue() *vertex {
	x := (*q).data[(*q).head]
	(*q).head++
	if (*q).head == (*q).cap { (*q).head = 0 }
	(*q).count--
	return x
}

type vertex struct {
	mark bool
	parent *vertex
	comp int
	edges []int
}


func DFS1(G *[]vertex, q *queue, bridge *int)  {
	for i, v := range *G {
		if !v.mark {
			v.parent = nil
			*bridge--
			VisitVertex1(G, q, &(*G)[i])
		}
	}
}

func VisitVertex1(G *[]vertex, q *queue, vert *vertex)  {
	(*vert).mark = true
	(*q).enqueue(vert)
	for _, v := range (*vert).edges {
		if !(*G)[v].mark {
			(*G)[v].parent = vert
			VisitVertex1(G, q, &((*G)[v]))
		}
	}
}

func DFS2(G *[]vertex, q *queue, bridge *int)   {
	var v *vertex
	component := 0
	for !(*q).empty() {
		v = (*q).dequeue()
		if (*v).comp == -1 {
			*bridge++
			VisitVertex2(G, v, component)
			component++
		}
	}

}

func VisitVertex2(G *[]vertex, vert *vertex, component int)  {
	(*vert).comp = component
	for _, e := range (*vert).edges {
		if (*G)[e].comp == -1 && (*G)[e].parent != vert {
			VisitVertex2(G, &((*G)[e]), component)
		}
	}
}

func main()  {
	var n, m, v1, v2 , bridge int
	bridge = 0
	//_, _ = fmt.Scanf("%d\n%d", &n, &m)
	input.Scanf("%d\n%d", &n, &m)
	var q queue
	q.init(n)
	incident := make([]vertex, 0)
	for i := 0; i < n; i++ {
		var v vertex
		v.mark = false
		v.comp = -1
		v.edges = make([]int, 0)
		incident = append(incident, v)
	}
	for i := 0; i < m; i++ {
		//_, _ = fmt.Scanf("%d %d", &v1, &v2)
		input.Scanf("%d %d", &v1, &v2)
		incident[v1].edges = append(incident[v1].edges, v2)
		incident[v2].edges = append(incident[v2].edges, v1)
	}
	DFS1(&incident, &q, &bridge)
	DFS2(&incident, &q, &bridge)
	fmt.Println(bridge)
}
