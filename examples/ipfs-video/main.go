//A very simple boilerplate webserver to serve a single page.
//This is used to quickly test out clientside web stuff that doesn't require anything serverside.
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/gopherjs/gopherjs.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		http.ServeFile(w, r, "gopherjs/gopherjs.js.gz")
	})
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.png")
	})

	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Println(err)
	}
}
