package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200 OK. Serving: %s\n", r.URL.Path)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page. Serving: %s\n", r.URL.Path)
}

func main() {
	PORT := ":4040"
	r := http.NewServeMux()
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
