package main

import (
	"fmt"
	"net/http"
)

const port = ":3333"

var route = []string{"/hello"}

func main() {

	http.HandleFunc(route[0], func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello\n")
	})

	fmt.Printf("hello listening: %s\n", port)
	fmt.Printf("route path: %s\n", route[0])
	http.ListenAndServe(port, nil)
}
