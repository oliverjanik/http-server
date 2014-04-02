package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type loggingResponseWriter struct {
	http.ResponseWriter
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	log.Printf("%d", code)
	w.ResponseWriter.WriteHeader(code)
}

func wrapHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := &loggingResponseWriter{w}
		log.Printf("%s %s", r.Method, r.URL)
		handler.ServeHTTP(lw, r)
	})
}

var port = flag.Int("p", 8001, "Port to run on")

func main() {
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on port %d", *port)

	addr := fmt.Sprintf(":%d", *port)

	log.Fatal(http.ListenAndServe(addr, wrapHandler(http.FileServer(http.Dir(wd)))))
}
