package handlers

import (
	"encoding/json"
	"fizzbuzz_rest/services/fizzbuzz"
	"net/http"
)

// Fizzbuzz is the handler that serves fizzbuzz requests.
type Fizzbuzz struct {
}

// Response represents the response send to the client.
type Response struct {
	Data  []string `json:"data,omitempty"`
	Error string   `json:"error,omitempty"`
}

// ServeHTTP allows Fizzbuzz to implement http.Handler interface.
func (h Fizzbuzz) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Validating data.
	fizzbuzzer, err := fizzbuzz.ValidateDataFizzbuzz(r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := Response{
			Error: err.Error(),
		}
		buf, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
		w.Write(buf)
		return
	}

	// Running Fizzbuzz and checking for error
	result, err := fizzbuzzer.Fizzbuzz()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := Response{
			Error: err.Error(),
		}
		buf, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
		w.Write(buf)
		return
	}

	// Marshalling and sending.
	resp := Response{
		Data: result,
	}
	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
