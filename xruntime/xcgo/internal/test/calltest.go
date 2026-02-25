package xcgotest

/*
#include "calltest.h"
*/
import "C"
import (
	"testing"
	"unsafe"

	"go.jcbhmr.com/xstd/xruntime/xcgo"
)

func TestCall_getFortyTwo(t *testing.T) {
	cFn := unsafe.Pointer(C.get_forty_two)
	cResult, err := xcgo.Call(uintptr(cFn))
	if err != nil {
		t.Fatalf("Call failed: %v", err)
	}
	result := int(cResult)
	if result != 42 {
		t.Fatalf("want %d, got %d", 42, result)
	}
}

func TestCall_add(t *testing.T) {
	cFn := unsafe.Pointer(C.add)
	cResult, err := xcgo.Call(uintptr(cFn), uintptr(1), uintptr(2))
	if err != nil {
		t.Fatalf("Call failed: %v", err)
	}
	result := int(cResult)
	if result != 3 {
		t.Fatalf("want %d, got %d", 3, result)
	}
}
