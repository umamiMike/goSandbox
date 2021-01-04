package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	// Example of using a http.FileServer if you want "server-like" rather than "middleware" behavior
	// mux.Handle("/public", http.FileServer(http.Dir("/home/public")))

	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir("/home")))
	n.UseHandler(mux)

	http.ListenAndServe(":3002", n)
}
