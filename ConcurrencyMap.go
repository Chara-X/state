package state

// Only one goroutine can access the map, other goroutines through the channel to send the function to the goroutine to access the map.
type ConcurrencyMap[K comparable, V any] struct {
	m map[K]V
	c chan func()
}

func NewConcurrencyMap[K comparable, V any]() ConcurrencyMap[K, V] {
	var m = ConcurrencyMap[K, V]{make(map[K]V), make(chan func(), 1024)}
	go func() {
		for f := range m.c {
			f()
		}
	}()
	return m
}
func (m *ConcurrencyMap[K, V]) Get(key K) (V, bool) {
	var receiver = make(chan V)
	var ok bool
	m.c <- func() {
		var value V
		value, ok = m.m[key]
		receiver <- value
	}
	return <-receiver, ok
}
func (m *ConcurrencyMap[K, V]) Set(key K, value V) {
	m.c <- func() {
		m.m[key] = value
	}
}
func (m *ConcurrencyMap[K, V]) Delete(key K) {
	m.c <- func() {
		delete(m.m, key)
	}
}
func (m *ConcurrencyMap[K, V]) Close() {
	close(m.c)
}
