package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "Hi there, my name is %s!", r.FormValue("name"))
}

func main() {
	http.HandleFunc("/a-path", handler)
	http.ListenAndServe(":8000", nil)
}
