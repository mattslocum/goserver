package routes

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/mattslocum/goserver/internal/logger"
	"github.com/mattslocum/goserver/internal/memoryStore"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// singleton
var instance *HashRouter
var once sync.Once

/**
 * Construct a HashRouter. Using this to do our DI since we don't have an IOC framework
 */
func GetHashRouter(wg *sync.WaitGroup) *HashRouter {
	once.Do(func() {
		instance = &HashRouter{
			store: memoryStore.Cache,
			sleep: 5,
			wg: wg,
		}
	})
	return instance
}

type HashRouter struct {
	mu sync.Mutex // guard count
	count  int
	store memoryStore.ICacheStore
	sleep int // Used to show latency
	wg *sync.WaitGroup
}

// Main entry point for HashRouter.
// Handles both GET and POST
func (h *HashRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.get(w, req)
	case "POST":
		h.post(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not supported")
	}
}

func (h *HashRouter) get(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/hash/"))

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "invalid hash ID. Example: /hash/123")
		return
	}
	// we are okay if the number doesn't exist yet.

	fmt.Fprintf(w, "%s", h.store.Get(id))
}

func (h *HashRouter) post(w http.ResponseWriter, req *http.Request) {
	// validation
	password := req.FormValue("password")
	if password == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "'password' is required")
		return
	}

	num := h.inc()
	fmt.Fprintf(w, "%d", num)
	go h.hash(num, password)
}

func (h *HashRouter) inc() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.count++
	return h.count
}

// Not sure if this should be considered the biz logic or if it is expected that hash is 3rd party service.
func (h *HashRouter) hash(num int, password string) {
	h.wg.Add(1)
	time.Sleep(time.Duration(h.sleep) * time.Second) // Latency! yay!
	hasher := sha512.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	h.store.Put(num, sha)
	logger.Debug.Println("writing hash", sha)
	h.wg.Done()
}
