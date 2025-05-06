package main

import "fmt"

func main() {
	sl := make([]int, 0, 4)
	sl = append(sl, 1, 2, 3)
	sl = addToSlice(sl, 1)
	fmt.Println(sl)

	slp := make([]int, 0, 4)
	slp = append(slp, 1, 2, 3)
	addToSlicePtr(&slp, 1)
	fmt.Println(slp)
}

func addToSlice(sl []int, v int) []int {
	sl = append(sl, v)
	return sl
}

func addToSlicePtr(slp *[]int, v int) {
	*slp = append(*slp, v)
}
