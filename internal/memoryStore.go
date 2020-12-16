package memorystore

// ICacheStore interface so we could swap out with backing store later.
type ICacheStore interface {
	Get(key int) string
	Put(key int, val string)
}

// TODO: thread safe? But because of our implementation of /hash, we know the key is thread safe.
type MemoryStore struct {
	// Really wish Golang supported generics, but it looks like they will soon.
	data map[int]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[int]string)}
}

func (m *MemoryStore) Get(key int) string {
	return m.data[key]
}

func (m *MemoryStore) Put(key int, val string) {
	m.data[key] = val
}

// Singleton
var Cache = NewMemoryStore()

