package main

import (
	"container/heap"
	"fmt"
)

type vertex struct {
	index int
	dist int
	w int
	name point
	edges []point
}

type point struct {
	x int
	y int
}

type PriorityQueue []*vertex

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*vertex)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(elem *vertex) { heap.Fix(pq, elem.index) }

func (pq *PriorityQueue) decrease( index, len int) {
	i := index
	(*pq)[index].dist = len
	for i > 0 && (*pq)[(i-1)/2].dist > len {
		pq.Swap((i-1)/2, i)
		(*pq)[i].index = i
		i = (i - 1) / 2
	}
	(*pq)[i].index = i
}

func relax(u, v *vertex, w int) bool {
	changed := u.dist + w < v.dist
	if changed { v.dist = u.dist + w }
	return changed
}

func Dijkstra(G *[][]vertex, s int) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &(*G)[0][0])
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			if i == 0 && j == 0 {
				(*G)[i][j].dist = (*G)[i][j].w
			} else {
				(*G)[i][j].dist = 13501
			}
			heap.Push(&pq, &(*G)[i][j])
		}
	}
	for pq.Len() > 0{
		v := heap.Pop(&pq).(*vertex)
		for _, e := range v.edges {
			if (*G)[e.x][e.y].index != -1 && relax(v, &(*G)[e.x][e.y], (*G)[e.x][e.y].w) {
				pq.decrease((*G)[e.x][e.y].index, (*G)[e.x][e.y].dist)
			}
		}
	}
}

func main() {
	var n, w int
	_, _ = fmt.Scanf("%d", &n)
	incident := make([][]vertex, n)
	for i := 0; i < n; i++ {
		incident[i] = make([]vertex, n)
		for j := 0; j < n; j++ {
			incident[i][j].edges = make([]point, 0)
		}
	}
	var p point
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&w)
			p.x = i
			p.y = j
			incident[i][j].name = p
			incident[i][j].w = w
			if j > 0 { incident[i][j - 1].edges = append(incident[i][j - 1].edges, p) }
			if j < n - 1 { incident[i][j + 1].edges = append(incident[i][j + 1].edges, p) }
			if i > 0 { incident[i - 1][j].edges = append(incident[i - 1][j].edges, p) }
			if i < n - 1 { incident[i + 1][j].edges = append(incident[i + 1][j].edges, p) }
		}
	}
	Dijkstra(&incident, n)
	fmt.Println(incident[n - 1][n - 1].dist)
}