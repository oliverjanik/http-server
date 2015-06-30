package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func runProxy(server string, localPort int) {
	addr := fmt.Sprintf(":%v", localPort)

	log.Println("Proxying on", addr, "to", server)

	client := &http.Client{}

	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		url := "http://" + server + "/" + r.RequestURI

		req, err := http.NewRequest(r.Method, url, r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(502)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(502)
			return
		}

		for key, val := range resp.Header {
			w.Header().Del(key)
			for _, v := range val {
				w.Header().Add(key, v)
			}
		}

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			log.Println(err)
		}

		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	})))
}
