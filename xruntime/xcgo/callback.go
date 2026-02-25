package xcgo

/*
#include "callback.h"
#include <stdint.h>
#include <stddef.h>
*/
import "C"
import (
	"math"
	"reflect"
	"sync"
	"unsafe"
)

var callbacks = struct {
	sync.Mutex
	NextID int
	Value  map[uintptr]any
}{
	NextID: 0,
	Value:  map[uintptr]any{},
}

//export call_callback
func call_callback(caller C.callback_fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15 C.uintptr_t) C.uintptr_t {
	callbacks.Lock()
	defer callbacks.Unlock()

	fn := callbacks.Value[uintptr(unsafe.Pointer(caller))]
	fnv := reflect.ValueOf(fn)
	fnt := fnv.Type()
	cArgs := []C.uintptr_t{a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15}

	argsv := make([]reflect.Value, fnt.NumIn())
	for i := range argsv {
		in := fnt.In(i)
		var arg reflect.Value
		switch in.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Uintptr:
			arg = reflect.ValueOf(uintptr(cArgs[i])).Convert(in)
		case reflect.Float32:
			arg = reflect.ValueOf(math.Float32frombits(uint32(cArgs[i])))
		case reflect.Float64:
			arg = reflect.ValueOf(math.Float64frombits(uint64(cArgs[i])))
		case reflect.Bool:
			arg = reflect.ValueOf(cArgs[i] != 0)
		case reflect.Ptr, reflect.UnsafePointer:
			arg = reflect.NewAt(in, unsafe.Pointer(uintptr(cArgs[i]))).Elem()
		default:
			panic("unsupported argument type")
		}
		argsv[i] = arg
	}
	resultsv := fnv.Call(argsv)
	switch fnt.NumOut() {
	case 0:
		return 0
	case 1:
		resultv := resultsv[0]
		switch resultv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Uintptr:
			return C.uintptr_t(resultv.Convert(reflect.TypeOf(uintptr(0))).Uint())
		case reflect.Float32:
			return C.uintptr_t(math.Float32bits(float32(resultv.Float())))
		case reflect.Float64:
			return C.uintptr_t(math.Float64bits(resultv.Float()))
		case reflect.Bool:
			if resultv.Bool() {
				return 1
			}
			return 0
		case reflect.Pointer, reflect.UnsafePointer:
			return C.uintptr_t(uintptr(resultv.UnsafePointer()))
		default:
			panic("unsupported return type")
		}
	default:
		panic("multiple return values not supported")
	}
}

func NewCallback(fn any) uintptr {
	callbacks.Lock()
	defer callbacks.Unlock()

	fnv := reflect.ValueOf(fn)
	fnt := fnv.Type()
	if fnt.Kind() != reflect.Func {
		panic("fn must be a function")
	}
	if fnt.NumOut() > 1 {
		panic("fn must have at most one return value")
	}
	if fnt.NumIn() > 15 {
		panic("fn must have at most 15 parameters")
	}
	for i := range fnt.NumIn() {
		in := fnt.In(i)
		switch in.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Uintptr,
			reflect.Float32, reflect.Float64,
			reflect.Bool,
			reflect.Ptr, reflect.UnsafePointer:
			// supported
		default:
			panic("unsupported parameter type")
		}
	}
	if fnt.NumOut() == 1 {
		out := fnt.Out(0)
		switch out.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Uintptr,
			reflect.Float32, reflect.Float64,
			reflect.Bool,
			reflect.Pointer, reflect.UnsafePointer:
			// supported
		default:
			panic("unsupported return type")
		}
	}

	id := callbacks.NextID
	if id >= 2000 {
		panic("too many callbacks")
	}
	callbacks.NextID++

	ptr := uintptr(unsafe.Pointer(C.callbacks[id]))
	callbacks.Value[ptr] = fn
	return ptr
}
