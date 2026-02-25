package xlist

import (
	"container/list"
	"iter"
	"structs"
	"unsafe"
)

type List[T any] list.List

func New[T any]() *List[T] {
	return (*List[T])(list.New())
}

func (l *List[T]) Back() *Element[T] {
	return (*Element[T])(unsafe.Pointer((*list.List)(l).Back()))
}

func (l *List[T]) Front() *Element[T] {
	return (*Element[T])(unsafe.Pointer((*list.List)(l).Front()))
}

func (l *List[T]) InsertBefore(value T, mark *Element[T]) *Element[T] {
	return (*Element[T])(unsafe.Pointer((*list.List)(l).InsertBefore(value, &mark.inner)))
}

func (l *List[T]) InsertAfter(value T, mark *Element[T]) *Element[T] {
	return (*Element[T])(unsafe.Pointer((*list.List)(l).InsertAfter(value, &mark.inner)))
}

func (l *List[T]) Len() int {
	return (*list.List)(l).Len()
}

func (l *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := l.Front(); e != nil; e = e.Next() {
			if !yield(e.Value()) {
				return
			}
		}
	}
}

type Element[T any] struct {
	_     structs.HostLayout
	inner list.Element
}

func (e *Element[T]) Value() T {
	return e.inner.Value.(T)
}

func (e *Element[T]) Next() *Element[T] {
	return (*Element[T])(unsafe.Pointer(e.inner.Next()))
}

func (e *Element[T]) Prev() *Element[T] {
	return (*Element[T])(unsafe.Pointer(e.inner.Prev()))
}
