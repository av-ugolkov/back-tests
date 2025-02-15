package sum

//#cgo LDFLAGS: -L. -lstdc++
//#cgo CXXFLAGS: -std=c++17 -I.
//#include "wrapper.h"
import "C"
import "unsafe"

type CppSum struct {
	w C.SumWrapper
}

func New() *CppSum {
	return &CppSum{w: C.Init()}
}

func (s CppSum) Destroy() {
	C.Destroy((C.SumWrapper)(unsafe.Pointer(s.w)))
}

func (s CppSum) Sum(a, b int) int {
	return int(C.Sum((C.SumWrapper)(unsafe.Pointer(s.w)), C.int(a), C.int(b)))
}
