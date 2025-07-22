package main

import (
	"fmt"
)

func main() {
	arr := []string{"str0", "str1", "str2", "str3", "str4", "str5"}

	baseStr := arr[0]
	for _, s := range arr[1:] {
		baseStr = joinString(baseStr, s)
	}

	fmt.Println(baseStr)
}

func joinString(base, str string) string {
	return fmt.Sprintf("%s;%s", base, str)
}
