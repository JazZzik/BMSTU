package main

import "fmt"

var a = make([]int , 0)

func partition(low, high int, less func(i, j int) bool, swap func(i, j int)) int{
	i := low
	j := low
	for j < high {
		if less(j, high) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, high)
	return i
}

func qsortrec(low, high int, less func(i, j int) bool, swap func(i, j int)){
	if low < high {
		q := partition(low, high, less, swap)
		qsortrec(low, q - 1, less, swap)
		qsortrec(q + 1, high, less, swap)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)){
	qsortrec(0, n - 1, less, swap)
}


func main() {
	var n, m, i int
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&m)
		a = append(a, m)
	}
	qsort(n, func (i , j int ) bool { return a[i] < a[j] }, func (i , j int ) { a[i], a[j] = a[j], a[i] })
	for i := 0; i < n; i++ {
	        fmt.Printf("%d ", a[i])
	}
}