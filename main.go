package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Create a map to hold the response data
		response := map[string]string{
			"message": "Hello World",
			"time":    time.Now().Format(time.RFC850),
			"host":    r.Host,
		}

		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":9990", nil)
}
