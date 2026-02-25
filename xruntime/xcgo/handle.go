//go:build cgo

package xcgo

import "runtime/cgo"

type Handle[T any] cgo.Handle

func NewHandle[T any](v T) Handle[T] {
	return Handle[T](cgo.NewHandle(v))
}

func (h Handle[T]) Value() T {
	return cgo.Handle(h).Value().(T)
}

func (h Handle[T]) Delete() {
	_ = cgo.Handle(h).Value().(T)
	cgo.Handle(h).Delete()
}
