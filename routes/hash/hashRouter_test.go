package routes

import (
	memorystore "github.com/mattslocum/goserver/internal"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestHashRouter_ServeHTTP_post_inc(t *testing.T) {
	router, _ := newRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	for i := 1; i <= 5; i++ {
		res, err := http.PostForm(ts.URL, url.Values{"password": {"hunter2"}})
		if err != nil {
			t.Error(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Error(err)
		}
		if string(body) != strconv.Itoa(i) {
			t.Errorf("Wrong POST Response. got: %s, want: %d", body, i)
		}
	}
}

func TestHashRouter_ServeHTTP_post_hash(t *testing.T) {
	router, store := newRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	tests := []struct{password string; hash string}{
		{
			password: "password",
			hash:     "sQnzu7wkTrgkQZF-0G1hi5AI3Qmzvv0bXgc5THBqi7mAsdd4Xll27ASbRt9fEyavWi6m0QP9B8lThf-rDKy8hg==",
		},
		{
			password: "hunter2",
			hash:     "a5ftaNFOs_GqlZzl1Jx9xhLh6x2v1zsecFhHSD_WpsgJ8s606N9v-ZhMYpj_AoXKzmYUv42qnwBwEBtsiYmeIg==",
		},
		{
			password: "angryMonkey",
			hash:     "ZEHhWB65gUlzdVwtDQArEyx-KVLzp_aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A-gf7Q==",
		},
	}

	for i, test := range(tests) {
		_, err := http.PostForm(ts.URL, url.Values{"password": {test.password}})
		if err != nil {
			t.Error(err)
		}

		// HashDone feels a little hacky, but this is contrived anyway.
		sha := store.Get(i + 1) // since we output starting at 1
		if sha != test.hash {
			t.Errorf("Wrong Hash. got: %s, want: %s", sha, test.hash)
		}
	}
}

func TestHashRouter_ServeHTTP_get(t *testing.T) {
	router, store := newRouter(memorystore.NewMemoryStore())
	ts := httptest.NewServer(router)
	defer ts.Close()

	store.Put(2, "two")

	res, err := http.Get(ts.URL + "/hash/2")
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if string(body) != "two" {
		t.Errorf("Wrong POST Response. got: %s, want: %s", body, "two")
	}
}


func newRouter(stores ...*memorystore.MemoryStore) (*HashRouter, *memorystore.MemoryStore) {
	var store *memorystore.MemoryStore
	if len(stores) > 0 {
		store = stores[0]
	} else {
		store = memorystore.NewMemoryStore()
	}
	return &HashRouter{store: store, sleep: 0, HashDone: make(chan string)}, store
}
