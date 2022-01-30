package main

import (
	"fmt"
	"math/big"
)

var o = big.NewRat(0, 1)
var m[][]big.Rat

func subtract(n, line1, line2 int) { // k = k - l
	for j := 0; j < n+1; j++ {
		m[line1][j].Sub(&m[line1][j], &m[line2][j])
	}
}

func refactor(lineNum, n int) {
	if m[lineNum][lineNum].Cmp(o) != 0 {
		return
	}
	for i := lineNum + 1; i < n; i++ {
		if m[i][lineNum].Cmp(o) != 0 {
			m[lineNum], m[i] = m[i], m[lineNum]
			return
		}
	}
}


func main() {
	var n int
	var x int64
	_, _ = fmt.Scan(&n)
	m = make([][]big.Rat, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ {
			_, _ = fmt.Scan(&x)
			k := big.NewRat(x, 1)
			m[i] = append(m[i], *k)
		}
	}

	for i := 0; i < n - 1; i++ {
		var num = big.Rat{}
		for j := i; j < n; j++ {
			if m[j][i].Cmp(o) != 0 {
				num.Inv(&m[j][i])
				for k := i; k < n + 1; k++ {
					if m[j][k].Cmp(o) != 0 {
						m[j][k].Mul(&m[j][k], &num)
					}
				}
			}
		}
		refactor(i, n)
		for j := i + 1; j < n; j++ {
			if m[j][i].Cmp(o) != 0 {
				subtract(n, j, i)
			}
		}
	}

	for i := 0;i < n; i++{
		if m[i][i].Cmp(o) == 0 {
			fmt.Print("No solution")
			return
		}
	}
	var t = big.Rat{}
	res := make([]big.Rat, n)
	for i := n - 1; i > -1; i--{
		tmp := m[i][n]
		for j := n - 1; j > i; j--{
			t.Mul(&m[i][j], &res[j])
			tmp.Sub(&tmp, &t)
		}
		t.Inv(&m[i][i])
		res[i].Mul(&tmp, &t)
	}
	for i := 0; i < n; i++{
		fmt.Println(res[i].String())
	}
}
