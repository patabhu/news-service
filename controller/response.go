package controller

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, r interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(r)
	if err != nil {
		return
	}
	_, _ = w.Write(b)
}
