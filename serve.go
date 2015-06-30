package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

type loggingResponseWriter struct {
	http.ResponseWriter
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	log.Printf("%v", code)
	w.ResponseWriter.WriteHeader(code)
}

func wrapHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := &loggingResponseWriter{w}
		log.Printf("%s %s", r.Method, r.URL)
		handler.ServeHTTP(lw, r)
	})
}

func runServe(dir string, port int) {
	addr := fmt.Sprintf(":%v", port)

	absPath, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}

	log.Println("Listening on port", port)
	log.Println("Serving", absPath)

	defer openBrowser(addr)

	log.Fatal(http.ListenAndServe(addr, wrapHandler(http.FileServer(http.Dir(absPath)))))
}

func openBrowser(url string) {
	url = "http://" + url

	log.Printf("Opening %v", url)

	exec.Command("cmd", "/c", "start", url).Start()
}
