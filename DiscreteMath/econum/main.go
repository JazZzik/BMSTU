package main

import ("fmt"
	"strings")

func main()  {
	var str string
	fmt.Scan(&str)
	var i, j, count int
	j = strings.Index(str, ")")
	for j != -1 {
		i = strings.LastIndex(str[: j], "(")
		j++
		str = strings.ReplaceAll(str, str[i : j], "")
		count++
		j = strings.Index(str, ")")
	}
	fmt.Println(count)
}
