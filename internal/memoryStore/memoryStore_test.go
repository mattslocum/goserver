package memoryStore

import "testing"

func TestMemoryStore_Put(t *testing.T) {
	store := &MemoryStore{data: make(map[int]string)}
	store.Put(5, "five")

	result := store.Get(5)
	if result != "five" {
		t.Errorf("Put(%d) was incorrect, got: %s, want: %s", 5, result, "five")
	}
}

func TestMemoryStore_Get(t *testing.T) {
	store := &MemoryStore{data: make(map[int]string)}
	store.Put(6, "six")

	result := store.Get(6)
	if result != "six" {
		t.Errorf("Put(%d) was incorrect, got: %s, want: %s", 5, result, "six")
	}

	// should still be there.
	result2 := store.Get(6)
	if result2 != "six" {
		t.Errorf("Put(%d) was incorrect, got: %s, want: %s", 5, result2, "six")
	}
}
