package routes

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var cache = make(map[int]string)


type HashHandler struct {
	mu sync.Mutex // guards count
	count  int
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
	fmt.Fprintf(w, "%s", cache[id])
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
	cache[num] = sha
	fmt.Println(num, sha)
}
