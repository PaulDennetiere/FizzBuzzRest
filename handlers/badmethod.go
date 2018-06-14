package handlers

import (
	"encoding/json"
	"net/http"
)

// BadMethod handles every bad method requests.
type BadMethod struct {
}

func (bm BadMethod) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Error: "Bad method",
	}

	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(buf)

}
