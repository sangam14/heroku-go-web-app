package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "gopherlabs")
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

