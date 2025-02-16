package main

import "fmt"

func main() {
	sl := make([]int, 0, 4)
	sl = append(sl, 1, 2, 3)
	addToSlice(sl)
	fmt.Println(sl)
}

func addToSlice(sl []int) []int {
	sl = append(sl, 4)
	return sl
}

func addToSlicePtr(slp *[]int) {
	*slp = append(*slp, 4)
}
