package main

import (
	"net/http"
	"strconv"
)

func main() {
	okBuf := []byte("OK")
	okLen := strconv.Itoa(len(okBuf))

	http.HandleFunc("/healthz", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
		rw.Header().Set("Content-Length", okLen)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(okBuf)
	})
	http.ListenAndServe(":4000", nil)
}
