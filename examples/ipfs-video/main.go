//A very simple boilerplate webserver to serve a single page.
//This is used to quickly test out clientside web stuff that doesn't require anything serverside.
package main

import (
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	jsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "gopherjs/gopherjs.js")
	})
	http.Handle("/gopherjs/gopherjs.js", gziphandler.GzipHandler(jsHandler))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.png")
	})

	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Println(err)
	}
}
