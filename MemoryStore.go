package state

import (
	"time"
)

// Every key-value pair in the store has a timer that will be reset when the key is accessed. If the key is not accessed for the duration, the key-value pair will be deleted by the timer's callback function.
type MemoryStore[T any] struct {
	store ConcurrencyMap[string, entry[T]]
}

func NewMemoryStore[T any]() MemoryStore[T] {
	return MemoryStore[T]{store: NewConcurrencyMap[string, entry[T]]()}
}
func (store *MemoryStore[T]) Get(key string) (T, bool) {
	if entry, ok := store.store.Get(key); ok {
		entry.timer.Reset(entry.duration)
		return entry.value, true
	}
	return *new(T), false
}
func (store *MemoryStore[T]) Set(key string, value T, duration time.Duration) {
	if entry, ok := store.store.Get(key); ok {
		entry.timer.Stop()
	}
	store.store.Set(key, entry[T]{value, time.AfterFunc(duration, func() { store.store.Delete(key) }), duration})
}
func (store *MemoryStore[T]) Close() {
	store.store.Close()
}

type entry[T any] struct {
	value    T
	timer    *time.Timer
	duration time.Duration
}
