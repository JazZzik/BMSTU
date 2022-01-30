package main

import (
	"fmt"
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
	name int
	edges []int
}


var (
	dist [][]int
)

func BFS(incident []vertex, vert int, j int)  {
	for i := range incident { incident[i].mark = false }
	var q queue
	q.init(len(incident))
	incident[vert].mark = true
	q.enqueue(&incident[vert])
	for !q.empty() {
		u := q.dequeue()
		for _, v := range (*u).edges {
			if !incident[v].mark {
				dist[j][v] =  dist[j][u.name] + 1
				incident[v].mark = true
				q.enqueue(&incident[v])
			}
		}
	}
}

func main() {
	var n, m, v1, v2, k int
	_, _ = fmt.Scanf("%d\n%d", &n, &m)
	incident := make([]vertex, 0)
	for i := 0; i < n; i++ {
		var v vertex
		v.name = i
		v.edges = make([]int, 0)
		incident = append(incident, v)
	}
	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d %d", &v1, &v2)
		incident[v1].edges = append(incident[v1].edges, v2)
		incident[v2].edges = append(incident[v2].edges, v1)
	}
	_, _ = fmt.Scan(&k)
	dist = make([][]int, k)
	for i := 0; i < k; i++ {
		_, _ = fmt.Scanf("%d", &v1)
		dist[i] = make([]int, n)
		BFS(incident, v1, i)
	}
	EqDist := make([]int, 0)
	for i := 0; i < n; i++ {
		f := true
		for j := 0; j < k - 1 && f; j++ {
			if dist[j][i] != dist[j + 1][i] || dist[j][i] == 0 { f = false }
		}
		if f { EqDist = append(EqDist, i) }
	}

	if len(EqDist) == 0 { fmt.Println("-")
	} else { for _, v := range EqDist { fmt.Printf("%d ", v) } }
}
