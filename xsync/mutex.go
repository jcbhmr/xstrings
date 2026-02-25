package xsync

import "sync"

type Mutex[T any] struct {
	mu    sync.Mutex
	value T
}

func (m *Mutex[T]) Lock() (value *T, unlock func()) {
	m.mu.Lock()
	return &m.value, m.mu.Unlock
}

func (m *Mutex[T]) TryLock() (value *T, unlock func(), ok bool) {
	if !m.mu.TryLock() {
		return nil, nil, false
	}
	return &m.value, m.mu.Unlock, true
}
