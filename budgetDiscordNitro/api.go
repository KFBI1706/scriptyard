package main

import (
	"encoding/json"
	"net/http"
)

func startAPI() {
	http.HandleFunc("/api/v1/emojis", func(w http.ResponseWriter, r *http.Request) {
		emojis.RLock()

		b, err := json.Marshal(emojis.List)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":1337", nil)

}
