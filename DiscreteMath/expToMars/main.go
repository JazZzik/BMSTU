package main
// rot ebal
import "fmt"

var compat [][]int
var group1 []int
var group2 []int

func formgroup(arr [][]string, dop []int, k0, k1, k2, pos, n int) {
	fmt.Print(dop, k0, k1, k2, pos, "\n")
	if pos == n {
		compat = append(compat, dop)
		group1 = append(group1, k0)
		group2 = append(group2, k1)
		return
	}
	var i int
	for i = 0; i < n; i++ {
		if arr[pos][i] == "+" {
			break
		}
	}
	if i == n {
		dop[pos] = 0
		formgroup(arr, dop, k0 + 1, k1, k2, pos + 1, n)
	} else {
		tmp1 := make([]int, n)
		tmp2 := make([]int, n)
		copy(tmp1, dop)
		copy(tmp2, dop)
		tmp1[pos], tmp2[pos] = 1, 2
		f1, f2 := true, true
		for ; i < n && (f1 || f2); i++ {
			if arr[pos][i] == "+" {
				if dop[i] == 0 {
					tmp1[i] = 2
					tmp2[i] = 1
				}
				if dop[i] == 1 {
					f1 = false
				}
				if dop[i] == 2 {
					f2 = false
				}
			}
		}
		if f1 { formgroup(arr, tmp1, k0, k1 + 1, k2, pos + 1, n) }
		if f2 { formgroup(arr, tmp2, k0, k1, k2 + 1, pos + 1, n) }
	}
}

func main() {
	var i, j, n int
	var c string
	_, _ = fmt.Scan(&n)
	var arr [][]string
	dop := make([]int, n)
	for i = 0; i < n; i++ {
		arr = append(arr, []string{})
		for j = 0; j < n; j++ {
			_, _ = fmt.Scan(&c)
			arr[i] = append(arr[i], c)
		}
	}
	formgroup(arr, dop, 0, 0, 0, 0, n)
	if len(compat) == 0 {
		fmt.Printf("No solution")
	} else {
		for i = 0; i < len(compat); i++ { group1[i] = n / 2 - group2[i] }
		str := make([]string, len(compat))
		for i = 0; i < n; i++ {
			for j = 0; j < len(compat); j++ {
				if compat[j][i] == 0 && group1[j] > 0 {
					compat[j][i] = 1
					group1[j]--
				}
				if compat[j][i] == 0 && group1[j] == 0 { compat[j][i] = 2 }
				str[j] += fmt.Sprintf("%d", compat[j][i])
			}
		}
		min := str[0]
		for i = 1; i < len(compat); i++ { if min > str[i] { min = str[i] } }
		for i = 0; i < len(min); i++ {
			if min[i] == '1' {
				fmt.Printf("%d ", i + 1)
			}
		}
	}
}