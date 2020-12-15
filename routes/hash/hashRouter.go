package routes

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/mattslocum/goserver/internal"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// singleton
var instance *HashRouter
var mu = &sync.Mutex{}

/**
 * Construct a HashRouter. Using this to do our DI since we don't have an IOC framework
 */
func NewHashRouter() *HashRouter {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &HashRouter{store: memorystore.Cache}
	}
	return instance
}

type HashRouter struct {
	mu sync.Mutex // guards count
	count  int
	store memorystore.ICacheStore
}

func (h *HashRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.get(w, req)
	case "POST":
		h.post(w, req)
	}
}

func (h *HashRouter) get(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/hash/"))
	fmt.Fprintf(w, "%s", h.store.Get(id))
}

func (h *HashRouter) post(w http.ResponseWriter, req *http.Request) {
	num := h.inc()
	fmt.Fprintf(w, "%d\n", num)

	password := req.FormValue("password")
	go h.hash(num, password)
}

func (h *HashRouter) inc() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.count++
	return h.count
}

func (h *HashRouter) hash(num int, password string) {
	time.Sleep(5 * time.Second) // Latency! yay!
	hasher := sha512.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	h.store.Put(num, sha)
}
