package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request in")
		fmt.Fprint(w, "hihi\n")
	})
	fmt.Println("listening: 3334")
	http.ListenAndServe(":3334", nil)
}
