package main

import (
	"go.guoyk.net/env"
	"nautilus/pkg/exe"
	"net/http"
	"strconv"
)

var (
	optTarget string
)

func main() {
	var err error
	defer exe.Exit(&err)

	if err = env.StringVar(&optTarget, "TARGET", ""); err != nil {
		return
	}

	okBuf := []byte("OK")
	okLen := strconv.Itoa(len(okBuf))

	http.HandleFunc("/healthz", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
		rw.Header().Set("Content-Length", okLen)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(okBuf)
	})
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		http.Redirect(rw, req, optTarget, http.StatusTemporaryRedirect)
	})
	http.ListenAndServe(":4000", nil)
}
