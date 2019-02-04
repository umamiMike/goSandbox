package main

import (
	"encoding/json"
	"net/http"
)

type Workout struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Musclegroup []string `json:"musclegroup"`
}
type jsonresp struct {
	w http.ResponseWriter
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8090", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	m := Workout{
		Name:        "testola",
		Description: "A Test Description",
		Musclegroup: []string{"Back", "legs"},
	}

	resp, _ := json.Marshal(m)

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	setAllHeaders(w)
	w.Write(resp)
}

func setAllHeaders(w http.ResponseWriter) {

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

}
