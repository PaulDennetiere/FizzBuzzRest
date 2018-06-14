package handlers

import (
	"encoding/json"
	"net/http"
)

// NotFound handles every request with a bad route.
type NotFound struct {
}

func (nf NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Error: "Bad route",
	}

	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(buf)
}
