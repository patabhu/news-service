package controller

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, r interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	b, err := json.Marshal(r)
	if err != nil {
		return
	}
	_, _ = w.Write(b)
}
