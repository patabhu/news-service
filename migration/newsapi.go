package migration

import (
	"encoding/json"
	"io"
	"net/http"
)

func getAPINewsDump(apiUrl string) *newsData {
	resp, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	d := newsData{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		panic(err)
	}
	return &d
}
