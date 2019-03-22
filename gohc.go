package gohc

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HealthCheck struct {
	condition func() bool
}

func MakeHealtCheck(condition func() bool) *HealthCheck {
	return &HealthCheck{
		condition: condition,
	}
}

func (hc *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cond := hc.condition()
	if cond {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("unhealthy"))
	} else {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("OK"))
	}
}

func (hc *HealthCheck) Listen(addr string) error {
	m := mux.NewRouter()
	m.HandleFunc("/", hc.ServeHTTP)
	m.HandleFunc("/health-check", hc.ServeHTTP)
	s := &http.Server{
		Addr:           addr,
		Handler:        m,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
