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
var instance *HashHandler
var mu = &sync.Mutex{}

/**
 * Construct a HashHandler. Using this to do our DI since we don't have an IOC framework
 */
func NewHashHandler() *HashHandler {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &HashHandler{store: memorystore.Cache}
	}
	return instance
}

type HashHandler struct {
	mu sync.Mutex // guards count
	count  int
	store memorystore.ICacheStore
}

func (h *HashHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.get(w, req)
	case "POST":
		h.post(w, req)
	}
}

func (h *HashHandler) get(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/hash/"))
	fmt.Fprintf(w, "%s", h.store.Get(id))
}

func (h *HashHandler) post(w http.ResponseWriter, req *http.Request) {
	num := h.inc()
	fmt.Fprintf(w, "%d\n", num)

	password := req.FormValue("password")
	go h.hash(num, password)
}

func (h *HashHandler) inc() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.count++
	return h.count
}

func (h *HashHandler) hash(num int, password string) {
	time.Sleep(5 * time.Second) // Latency! yay!
	hasher := sha512.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	h.store.Put(num, sha)
}
