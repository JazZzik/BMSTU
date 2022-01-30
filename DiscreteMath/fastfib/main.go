package main

import (
	"fmt"
	"math/big"
)


func multiplyMatrices(m1, m2 [4]big.Int) [4]big.Int {
	var m [4]big.Int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				t := big.NewInt(0)
				t.Mul(&m1[i * 2 + k], &m2[k * 2 + j])
				m[i * 2 + j].Add(t, &m[i * 2 + j])
			}
		}
	}
	return m
}

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 || n == 2 {
		fmt.Print(1)
	} else {
		n--
		m1 := [4]big.Int{*big.NewInt(1), *big.NewInt(1), *big.NewInt(1), *big.NewInt(0)}
		m2 := [4]big.Int{*big.NewInt(1), *big.NewInt(0), *big.NewInt(0), *big.NewInt(1)}
		for n > 0 {
			if n & 1 == 1 { m2 = multiplyMatrices(m2, m1) }
			m1 = multiplyMatrices(m1, m1)
			n >>= 1
		}
		res := big.NewInt(0)
		res.Add(&m2[2], &m2[3])
		fmt.Println(res)
	}
}
