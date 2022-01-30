package main

import "fmt"


func main()  {
	var n int
	var dividers []int
	_, _ = fmt.Scan(&n)
	fmt.Printf("graph {\n")
	if n == 1 {
		fmt.Printf("\t1\n")
	} else {
		dividers = make([]int, 0)
		tmp := make([]int, 0)
		for i := 1; i*i <= n; i++{
			if n % i == 0{
				if i*i == n {
					tmp = append(tmp, i)
				} else {
					tmp = append(tmp, i)
					dividers = append(dividers, n / i)
				}
			}
		}
		l := len(tmp)
		for i := range tmp { dividers = append(dividers, tmp[l - i - 1])}
		l = len(dividers)
		for _, x := range dividers { fmt.Printf("\t%d\n", x)}
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				if dividers[i]%dividers[j] == 0 {
					for k := i + 1; k <= j; k++ {
						if k == j {
							fmt.Printf("\t%d--%d\n", dividers[i], dividers[j])
						} else {
							if dividers[k]%dividers[j] == 0 && dividers[i]%dividers[k] == 0 {
								break
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("}")
}