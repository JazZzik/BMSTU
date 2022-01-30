package main

import (
	"fmt"
	//"github.com/skorobogatov/input"
	"math"
	"sort"
)

var roads []road

type attraction struct {
	x, y, depth int
	parent *attraction
}

type road struct {
	length int
	parent1, parent2 *attraction
}

func find(a *attraction) *attraction {
	root := a
	for root.parent != nil { root = root.parent }
	return root
}

func union(r1, r2 *attraction)  {
	if r1.depth > r2.depth {
		r2.parent = r1
	} else if r1.depth == r2.depth {
		r1.depth++
		r2.parent = r1
	} else {
		r1.parent = r2
	}
}

func spanningTree(n int)  float64 {
	var l float64 = 0
	var parent1, parent2, r1 ,r2 *attraction
	for i := 0; i < n; i++ {
		parent1 = roads[i].parent1
		parent2 = roads[i].parent2
		r1, r2 = find(parent1), find(parent2)
		if r1 != r2 {
			l += math.Sqrt(float64(roads[i].length))
			union(r1, r2)
		}
	}
	return l
}

func main()  {
	var n, x, y int
	_, _ = fmt.Scanf("%d", &n)
	//input.Scanf("%d", &n)
	attractions := make([]attraction, n)
	rlen := n * ( n - 1) / 2
	roads = make([]road, rlen)
	var a attraction
	var r road
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d %d", &x, &y)
		//input.Scanf("%d %d", &x, &y)
		a.x, a.y = x, y
		attractions[i] = a
	}
	rcount := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			r.parent1 = &attractions[i]
			r.parent2 = &attractions[j]
			r.length = (attractions[i].x - attractions[j].x)*(attractions[i].x - attractions[j].x) +
				(attractions[i].y - attractions[j].y)*(attractions[i].y - attractions[j].y)
			roads[rcount] = r
			rcount++
		}
	}
	roads = roads[:rcount]
	sort.Slice(roads, func (i , j int ) bool { return roads[i].length < roads[j].length })
	fmt.Printf("%.2f\n", spanningTree(rlen))
}
