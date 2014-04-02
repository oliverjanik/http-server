package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type loggingHandler struct {
	h http.Handler
}

func (f *loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL)
	f.h.ServeHTTP(w, r)
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	lh := &loggingHandler{http.FileServer(http.Dir(wd))}

	port := 8001

	log.Printf("Starting server on port %d", port)

	addr := fmt.Sprintf(":%d", port)

	log.Fatal(http.ListenAndServe(addr, lh))
}
