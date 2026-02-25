package xcgo

/*
#include "call.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

func Call(fn uintptr, args ...uintptr) (uintptr, error) {
	cFn := *(*unsafe.Pointer)(unsafe.Pointer(&fn))
	var cArgs *C.uintptr_t
	var cArgsLen C.size_t
	if len(args) > 0 {
		cArgs = (*C.uintptr_t)(unsafe.Pointer(&args[0]))
		cArgsLen = C.size_t(len(args))
	}
	cResult, err := C.call(cFn, cArgs, cArgsLen)
	runtime.KeepAlive(args)
	return uintptr(cResult), err
}
