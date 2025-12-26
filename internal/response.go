package internal

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message  string      `json:"message"`
	Articles interface{} `json:"articles,omitempty"`
}

func WriteResponse(w http.ResponseWriter, code int, r interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(r)
	if err != nil {
		return
	}
	_, _ = w.Write(b)
}
