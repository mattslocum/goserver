package routes

import (
	"github.com/mattslocum/goserver/internal/memoryStore"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"sync"
	"testing"
)

var wg sync.WaitGroup

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

		wg.Wait()
		sha := store.Get(i + 1) // since we output starting at 1
		if sha != test.hash {
			t.Errorf("Wrong Hash. got: %s, want: %s", sha, test.hash)
		}
	}
}

func TestHashRouter_ServeHTTP_post_invalid(t *testing.T) {
	router, _ := newRouter(memoryStore.NewMemoryStore())
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.PostForm(ts.URL, url.Values{"password": {}})
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	expect := "'password' is required"
	if string(body) != expect {
		t.Errorf("Wrong POST Response. got: %s, want: %s", body, expect)
	}
	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Wrong Status Code. got: %d, want: %d", res.StatusCode, http.StatusUnprocessableEntity)
	}
}

func TestHashRouter_ServeHTTP_get(t *testing.T) {
	router, store := newRouter(memoryStore.NewMemoryStore())
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

func TestHashRouter_ServeHTTP_get_invalid(t *testing.T) {
	router, _ := newRouter(memoryStore.NewMemoryStore())
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hash/hacker")
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	expect := "Invalid hash ID. Example: /hash/123"
	if string(body) != expect {
		t.Errorf("Wrong POST Response. got: %s, want: %s", body, expect)
	}
	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Wrong Status Code. got: %d, want: %d", res.StatusCode, http.StatusUnprocessableEntity)
	}
}


func newRouter(stores ...*memoryStore.MemoryStore) (*HashRouter, *memoryStore.MemoryStore) {
	var store *memoryStore.MemoryStore
	if len(stores) > 0 {
		store = stores[0]
	} else {
		store = memoryStore.NewMemoryStore()
	}
	return &HashRouter{store: store, sleep: 0, wg: &wg}, store
}
