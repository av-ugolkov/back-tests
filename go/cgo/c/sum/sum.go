package sum

//#cgo CFLAGS: -I./
//#cgo LDFLAGS: -L${SRCDIR}/ -lsum -Wl,-rpath=${SRCDIR}/
//#include "sum.h"
import "C"

func Sum(a, b int) int {
	return int(C.sum(C.int(a), C.int(b)))
}
