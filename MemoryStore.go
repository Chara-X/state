package state

import (
	"sync"
	"time"
)

type MemoryStore[T any] struct{ store *sync.Map }

func NewMemoryStore[T any]() MemoryStore[T] { return MemoryStore[T]{new(sync.Map)} }
func (s *MemoryStore[T]) Store(key string, value T, duration time.Duration) {
	s.Delete(key)
	s.store.Store(key, entry[T]{value, time.AfterFunc(duration, func() { s.store.Delete(key) }), duration})
}
func (s *MemoryStore[T]) Load(key string) (T, bool) {
	if v, ok := s.store.Load(key); ok {
		var entry = v.(entry[T])
		entry.timer.Reset(entry.duration)
		return entry.value, true
	}
	return *new(T), false
}
func (s *MemoryStore[T]) Delete(key string) {
	if v, ok := s.store.Load(key); ok {
		v.(entry[T]).timer.Stop()
	}
	s.store.Delete(key)
}

type entry[T any] struct {
	value    T
	timer    *time.Timer
	duration time.Duration
}
