package app

import "net/http"

func Run() error {
	http.HandleFunc("/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return err
	}

	return nil
}
