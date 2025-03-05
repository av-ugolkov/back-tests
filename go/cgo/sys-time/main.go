package main

/*
#include <sys/time.h>
*/
import "C"

import "fmt"

func main() {
	var tv C.struct_timeval
	C.gettimeofday(&tv, nil)
	fmt.Println("Секунды:", tv.tv_sec, "Микросекунды:", tv.tv_usec)
}
