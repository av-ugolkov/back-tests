package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "qwer;asdf;zxcv;wert;sdfg;xcvb;erty;dfgh;cvbn"
	sep := ";"

	res1 := split(str, sep)
	res2 := cut(str, sep)
	res3 := index(str, sep)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

func split(str, sep string) []string {
	return strings.Split(str, sep)
}

func cut(str, sep string) []string {
	count := strings.Count(str, sep)
	res := make([]string, 0, count+1)

	var before string
	b := true
	for b {
		before, str, b = strings.Cut(str, sep)
		res = append(res, before)
	}

	return res
}

func index(str, sep string) []string {
	count := strings.Count(str, sep)
	res := make([]string, 0, count+1)

	var ind int
	for {
		ind = strings.Index(str, sep)
		if ind < 0 {
			res = append(res, str)
			break
		}
		res = append(res, str[:ind])
		str = str[ind+1:]
	}

	return res
}
