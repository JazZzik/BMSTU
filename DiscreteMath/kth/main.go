package main

import (
	"fmt"
	"math"
)

func CountDigits (x uint64) uint64 {
	if x < 10 { return 1
	} else { return 1 + CountDigits(x / 10)}
}

func Pow10(x uint64) uint64{
	if x == 0 { return 1
	} else { return 10 * Pow10(x - 1) }
}

func GetIthDigit(x, i uint64) uint64 {
	return x / Pow10(CountDigits(x) - i) % 10
}

func CountDigitsUpToX(x uint64) uint64 {
	if x == 0 { return 0 }
	var radix, exponent, count, tmp uint64
	radix = 1
	exponent = 1
	for {
		tmp = exponent * 10 - exponent
		if x < tmp { break
		} else {
			count += tmp * radix
			x -= tmp
			radix++
			exponent *= 10
		}
	}
	if x != 0 {count += radix * x}
	return count
}

func GetDigit(p uint64) uint64{
	var l, r uint64
	r = uint64(math.Pow(2, 64))
	l = 0
	for r - l > 1{
		m := (l + r) >> 1
		if CountDigitsUpToX(m) >= p { r = m
		} else { l = m }
	}
	x := l
	if CountDigitsUpToX(l) <= p { x = r }
	p -= CountDigitsUpToX(x - 1)
	return GetIthDigit(x, p)
}

func main() {
	//var k uint64
	//fmt.Scan(&k)
	//k++
	var a = [...]uint64 {0, 1, 2, 8, 9, 16, 1000000000, 1000000000000, 102030405060708090}
	for i, s := range a{
		s++
		fmt.Println(i, s, GetDigit(s))
	}


}
