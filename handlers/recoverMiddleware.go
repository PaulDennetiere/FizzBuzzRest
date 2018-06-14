package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Recover can generate a middleware to handle panic in handlers. If a handler panics, a  response with a 500 internal error status code is returned.
// Such panic can happened on marshal.
type Recover struct {
	ErrorLogger *log.Logger
}

// GenerateMiddleware generates a middleware for recovering handlers' panics.
func (rh *Recover) GenerateMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				// Sending response to client.
				resp := Response{
					Error: "internal error",
				}
				buf, err := json.Marshal(resp)
				if err != nil {
					rh.ErrorLogger.Println("unable to marshal a respone while recovering from panic: " + err.Error())
				}
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(buf)

				if err, ok := r.(error); ok {
					rh.ErrorLogger.Println("Panic recovered: " + err.Error())
					return
				}
				// If the panic is unknown
				rh.ErrorLogger.Printf("Unknown panic cause: %v", r)
			}
		}()
		next.ServeHTTP(w, r)
	}
}
