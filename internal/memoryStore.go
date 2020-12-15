package memorystore

// TODO: ICacheStore interface here so we could swap out with backing store later.
type ICacheStore interface {
	Get(key int) string
	Put(key int, val string)
}

// TODO: thread safe? Because of other implementation of /hash, we know the key is thread safe though.
type MemoryStore struct {
	data map[int]string
}

func (m *MemoryStore) Get(key int) string {
	return m.data[key]
}

func (m *MemoryStore) Put(key int, val string) {
	m.data[key] = val
}

// Singleton
var Cache = &MemoryStore{data: make(map[int]string)}

