package main

import "fmt"

type edge struct {
	length int
	dest int
}

type vertex struct {
	name, index, key, value int
	edges []edge
}

type queue struct {
	heap []*vertex
	count int
}

func (q *queue) init(n int)  {
	q.heap = make([]*vertex, n)
	q.count = 0
}

func (q *queue) empty() bool{
	return q.count == 0
}

func (q *queue) heapify(i, n int)  {
	var l, j, r int
	for {
		l = 2 * i + 1
		r = l + 1
		j = i
		if l < n && q.heap[i].key > q.heap[l].key { i = l }
		if r < n && q.heap[i].key > q.heap[r].key { i = r }
		if i == j { break }
		q.heap[i], q.heap[j] = q.heap[j], q.heap[i]
		q.heap[i].index = i
		q.heap[j].index = j
	}
}

func (q *queue) insert(v *vertex)  {
	i := q.count
	q.count++
	q.heap[i] = v
	for i > 0 && q.heap[(i - 1) / 2].key > q.heap[i].key {
		q.heap[(i - 1) / 2], q.heap[i] = q.heap[i], q.heap[(i - 1) / 2]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index, v.index = i, i
}

func (q *queue) extractMin()  *vertex {
	min := q.heap[0]
	q.count--
	if !q.empty() {
		q.heap[0] = q.heap[q.count]
		q.heap[0].index = 0
		q.heapify(0, q.count)
	}
	return min
}

func (q *queue) decreaseKey(index, k int)  {
	i := index
	q.heap[index].key = k
	for i > 0 && q.heap[(i - 1) / 2].key > k {
		q.heap[(i - 1) / 2], q.heap[i] = q.heap[i], q.heap[(i - 1) / 2]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index = i
}

func MST_Prim(listIncidence []*vertex)  int {
	length := 0
	v := listIncidence[0]
	var q queue
	q.init(len(listIncidence) - 1)
	for {
		v.index = -2
		for _, e := range v.edges {
			if listIncidence[e.dest].index == -1 {
				listIncidence[e.dest].key = e.length
				listIncidence[e.dest].value = v.name
				q.insert(listIncidence[e.dest])
			} else if listIncidence[e.dest].index != -2 && e.length <= listIncidence[e.dest].key {
				listIncidence[e.dest].value = v.name
				q.decreaseKey(listIncidence[e.dest].index, e.length)
			}
		}
		if q.empty() { break }
		v = q.extractMin()
		length += v.key
	}
	return length
}

func main()  {
	var n,m, v1, v2, length int
	_, _ = fmt.Scanf("%d\n%d", &n, &m)
	listIncidence := make([]*vertex, 0)
	for i := 0; i < n; i++ {
		var v vertex
		v.name = i
		v.index = -1
		v.edges = make([]edge, 0)
		listIncidence = append(listIncidence, &v)
	}
	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d %d %d", &v1, &v2, &length)
		var e edge
		e.length = length
		e.dest = v2
		listIncidence[v1].edges = append(listIncidence[v1].edges, e)
		e.dest = v1
		listIncidence[v2].edges = append(listIncidence[v2].edges, e)
	}
	fmt.Println(MST_Prim(listIncidence))
}