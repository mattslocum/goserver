package routes

import (
	"github.com/mattslocum/goserver/internal/middleware"
	routes "github.com/mattslocum/goserver/routes/hash"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStatsRouter_ServeHTTP(t *testing.T) {
	router := &StatsRouter{}
	ts := httptest.NewServer(router)
	defer ts.Close()

	middleware.Timers[routes.HashTimerName] = &middleware.EventTimer{Count: 4, Duration:100}

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	expect := `{"total":4,"average":25}`
	if strings.TrimSpace(string(body)) != expect {
		t.Errorf("Wrong Response. got: %s, want: %s", body, expect)
	}
}

func TestStatsRouter_ServeHTTP_zero(t *testing.T) {
	router := &StatsRouter{}
	ts := httptest.NewServer(router)
	defer ts.Close()

	middleware.Timers[routes.HashTimerName] = &middleware.EventTimer{Count: 0, Duration:0}

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	expect := `{"total":0,"average":0}`
	if strings.TrimSpace(string(body)) != expect {
		t.Errorf("Wrong Response. got: %s, want: %s", body, expect)
	}
}
