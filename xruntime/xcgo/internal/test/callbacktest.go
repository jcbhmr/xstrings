package xcgotest

/*
#include "callbacktest.h"
#include <stdint.h>
*/
import "C"
import (
	"testing"
	"unsafe"

	"go.jcbhmr.com/xstd/xruntime/xcgo"
)

func TestNewCallback_add(t *testing.T) {
	cFn := xcgo.NewCallback(func(cA C.int) C.int {
		a := int(cA)
		result := a + 200
		return C.int(result)
	})
	cResult := C.call_adder_with_one((*[0]byte)(*(*unsafe.Pointer)(unsafe.Pointer(&cFn))))
	if int(cResult) != 201 {
		t.Errorf("want %d, got %d", 201, int(cResult))
	}
}
