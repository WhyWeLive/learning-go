package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/health/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
